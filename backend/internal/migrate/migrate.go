package main

import (
	"fmt"
	"log"

	psq "backend/internal/db"
	"backend/internal/models"
)

func main() {
	// Подключение к базе данных
	db, err := psq.ConnectToDB()

	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Создание расширения для UUID (если используется PostgreSQL)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Миграция всех таблиц
	err = db.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Company{},
	)

	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	fmt.Println("👍 Миграция успешно завершена")
	fmt.Println("✅ Тестовые данные успешно добавлены")
}
