package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) SendMessage(chatID int64, text string) error {
	if b.BotAPI == nil {
		return fmt.Errorf("bot API is not initialized")
	}

	msg := tgbotapi.NewMessage(chatID, text)
	_, err := b.BotAPI.Send(msg)
	if err != nil {
		log.Printf("Failed to send message to %d: %v", chatID, err)
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}
