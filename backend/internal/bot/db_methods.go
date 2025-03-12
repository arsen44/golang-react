package bot

import (
	"backend/internal/models"
	"context"
	"fmt"
)

func (b *Bot) findPersonByPhoneNumber(ctx context.Context, phoneNumber string) (int, error) {
	var user models.User

	// Use GORM to find the user by phone number
	result := b.db.Where("phone_number = ?", phoneNumber).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(user.ID), nil
}

// updateChatID updates the chat_id for a user using GORM
func (b *Bot) updateChatID(ctx context.Context, personID int, chatID int64) error {
	// Convert int64 to string for the ChatID field
	chatIDStr := fmt.Sprintf("%d", chatID)

	// Update the user's chat_id
	result := b.db.Model(&models.User{}).Where("id = ?", personID).Update("chat_id", chatIDStr)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// getVerificationCode gets the verification code for a user using GORM
func (b *Bot) getVerificationCode(ctx context.Context, personID int) (string, error) {
	var user models.User

	result := b.db.Select("verification_code").Where("id = ?", personID).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// Handle nil verification code
	if user.VerificationCode == nil {
		return "", nil
	}

	return *user.VerificationCode, nil
}
