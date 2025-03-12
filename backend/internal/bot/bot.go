package bot

import (
	"context"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

// TelegramToken is your bot token from BotFather

var (
	TelegramToken = os.Getenv("TelegramToken")
)

// Bot struct holds the database connection and bot instance
type Bot struct {
	db     *gorm.DB
	BotAPI *tgbotapi.BotAPI
}

// NewBot creates a new Bot instance
func NewBot(db *gorm.DB) BotInterface {
	return &Bot{
		db: db,
	}
}

// Run initializes and starts the Telegram bot
func (b *Bot) Run() error {
	var err error

	// Initialize Telegram bot
	b.BotAPI, err = tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	log.Printf("Authorized on account %s", b.BotAPI.Self.UserName)

	// Set up updates configuration
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get updates channel
	updates := b.BotAPI.GetUpdatesChan(u)

	// Process updates
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Handle different message types
		switch {
		case update.Message.Command() == "start":
			b.handleStart(update.Message)
		case update.Message.Contact != nil:
			b.handleContact(update.Message)
		default:
			// Handle other messages if needed
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я понимаю только команду /start или отправку контакта.")
			b.BotAPI.Send(msg)
		}
	}

	return nil
}

// handleStart handles the /start command
func (b *Bot) handleStart(message *tgbotapi.Message) {
	chatID := message.Chat.ID

	// Create keyboard with contact button
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonContact("Отправить номер телефона"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Для продолжения, 🔒 Нажмите на кнопку \"Передать номер телефона\". Если вы не видите эту кнопку, разверните меню, нажав на иконку в правом нижнем углу экрана, рядом с полем для ввода.")

	msg.ReplyMarkup = keyboard

	if _, err := b.BotAPI.Send(msg); err != nil {
		log.Printf("Error sending start message: %v", err)
	}
}

// handleContact processes the phone number contact
func (b *Bot) handleContact(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	phoneNumber := message.Contact.PhoneNumber

	// Check if the phone number belongs to a user
	userID, err := b.findPersonByPhoneNumber(context.Background(), phoneNumber)
	if err != nil {
		log.Printf("Error finding person: %v", err)
		b.BotAPI.Send(tgbotapi.NewMessage(chatID, "Произошла ошибка при поиске."))
		return
	}

	if userID == 0 {
		// If no user is found
		b.BotAPI.Send(tgbotapi.NewMessage(chatID, "Пользователь с данным номером телефона не зарегистрирован."))
		return
	}

	// Get verification code
	verificationCode, err := b.getVerificationCode(context.Background(), userID)
	if err != nil {
		log.Printf("Error getting verification code: %v", err)
		b.BotAPI.Send(tgbotapi.NewMessage(chatID, "Произошла ошибка при получении кода подтверждения."))
		return
	}

	// Update chat_id
	err = b.updateChatID(context.Background(), userID, chatID)
	if err != nil {
		log.Printf("Error updating chat_id: %v", err)
		b.BotAPI.Send(tgbotapi.NewMessage(chatID, "Не удалось сохранить chat_id."))
		return
	}

	// Create message with verification code
	confirmationText := fmt.Sprintf("🔑 Ваш код подтверждения: %s\n\nСпасибо за использование нашего сервиса! 😊", verificationCode)
	confirmationMsg := tgbotapi.NewMessage(chatID, confirmationText)

	// Send message
	if _, err := b.BotAPI.Send(confirmationMsg); err != nil {
		log.Printf("Error sending confirmation message: %v", err)
	}
}
