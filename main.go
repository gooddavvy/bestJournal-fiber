package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gooddavvy/bestJournal-fiber/handlers"
)

var (
	port = os.Getenv("PORT")
	app  = fiber.New()
)

func main() {
	// API routes
	app.Get("/api/journalPages/:id", handlers.JournalPagesAPI)
	app.Get("/api/*", handlers.MainAPI)

	// Creating a static logical code and listening to server
	app.Static("/", "./public")
	app.Listen(":" + port)
}
