package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	envViper "github.com/gooddavvy/bestJournal-fiber/env"
	otherFs "github.com/gooddavvy/bestJournal-fiber/fs"
	handlers "github.com/gooddavvy/bestJournal-fiber/handlers"
	vars "github.com/gooddavvy/bestJournal-fiber/var"
)

var (
	app = fiber.New()
)

func main() {
	_, err := otherFs.GetJson()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert port to string and enable envViber configuration
	envViper.Config(false)
	envViper.Load()
	vars.Port = envViper.Atos("PORT")

	// API routes
	app.Get("/api/journalPages", handlers.JournalPagesAPI)
	app.Get("/api/*", handlers.MainAPI)

	// Creating a static logical code and listening to server
	app.Static("/", "./public")
	app.Listen(":" + vars.Port)
}
