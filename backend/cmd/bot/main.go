package main

import (
	"log"

	"backend/internal/bot"
	"backend/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	// Initialize database connection
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the bot
	botHandler := bot.NewBot(database)

	// Start the bot
	if err := botHandler.Run(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
