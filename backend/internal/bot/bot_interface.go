package bot

type BotInterface interface {
	SendMessage(chatID int64, text string) error
	Run() error
}
