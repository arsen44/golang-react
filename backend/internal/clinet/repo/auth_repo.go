package repo

import (
	"backend/internal/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"backend/internal/bot"
	"backend/internal/utils"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type ClientRepo struct {
	db                   *gorm.DB
	botService           bot.BotInterface
	secret               []byte
	accessTokenDuration  time.Duration
	refreshTokenDuration time.Duration
}

func NewClientRepo(db *gorm.DB, botService bot.BotInterface) ClientRepoInterface {
	// Установка значений по умолчанию прямо в репозитории
	secret := []byte("your-secret-key")        // Замените на ваш секретный ключ
	accessTokenDuration := time.Hour * 24      // Например, 24 часа
	refreshTokenDuration := time.Hour * 24 * 7 // Например, 7 дней

	return &ClientRepo{
		db:                   db,
		botService:           botService,
		secret:               secret,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration}
}

func (a *ClientRepo) CreateClient(phoneNumber string) (string, string, error) {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var user models.User
	result := tx.Where("phone_number = ?", phoneNumber).First(&user)

	verificationCode := utils.GenerateVerificationCode()

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				Username:         phoneNumber,
				PhoneNumber:      phoneNumber,
				VerificationCode: &verificationCode,
			}
			if err := tx.Create(&user).Error; err != nil {
				tx.Rollback()
				return "", "", fmt.Errorf("failed to create user: %w", err)
			}

			client := models.Client{UserID: user.ID}
			if err := tx.Create(&client).Error; err != nil {
				tx.Rollback()
				return "", "", fmt.Errorf("failed to create client: %w", err)
			}
		} else {
			tx.Rollback()
			log.Printf("Database error: %v", result.Error)
			return "", "", fmt.Errorf("database operation failed: %w", result.Error)
		}
	} else {
		user.VerificationCode = &verificationCode
		if err := tx.Save(&user).Error; err != nil {
			tx.Rollback()
			return "", "", fmt.Errorf("failed to update verification code: %w", err)
		}

		var client models.Client
		clientResult := tx.Where("user_id = ?", user.ID).First(&client)

		if clientResult.Error == gorm.ErrRecordNotFound {
			client = models.Client{UserID: user.ID}
			if err := tx.Create(&client).Error; err != nil {
				tx.Rollback()
				return "", "", fmt.Errorf("failed to create client: %w", err)
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return "", "", fmt.Errorf("transaction commit failed: %w", err)
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      user.ID,
		"phone_number": user.PhoneNumber,
		"exp":          time.Now().Add(a.accessTokenDuration).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(a.refreshTokenDuration).Unix(),
		"type": "refresh",
	})

	signedAccessToken, err := accessToken.SignedString(a.secret)
	if err != nil {
		return "", "", fmt.Errorf("error while signing access token: %w", err)
	}

	signedRefreshToken, err := refreshToken.SignedString(a.secret)
	if err != nil {
		return "", "", fmt.Errorf("error while signing refresh token: %w", err)
	}

	if user.ChatID != nil && *user.ChatID != "" {
		chatIDInt, err := strconv.ParseInt(*user.ChatID, 10, 64)
		if err == nil {
			confirmationText := fmt.Sprintf("🔑 Ваш новый код подтверждения: %s\n\nСпасибо за использование нашего сервиса! 😊", verificationCode)
			log.Printf("Попытка отправки сообщения в чат %d: %s", chatIDInt, confirmationText)
			if err := a.botService.SendMessage(chatIDInt, confirmationText); err != nil {
				log.Printf("Ошибка при отправке сообщения: %v", err)
			}
		} else {
			log.Printf("Ошибка преобразования ChatID: %v", err)
		}
	}

	return signedAccessToken, signedRefreshToken, nil
}
