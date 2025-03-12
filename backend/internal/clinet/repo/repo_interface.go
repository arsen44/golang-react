package repo

// ClinetRepoInterface интерфейс для работы с созданием клиента
type ClientRepoInterface interface {
	CreateClient(phoneNumder string) (string, string, error)
}
