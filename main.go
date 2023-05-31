package main

import "github.com/gofiber/fiber/v2"

var (
	port = "2012"
	app  = fiber.New()
)

func main() {

	// API routes
	//app.Get("/api/*", func(c *fiber.Ctx) error {
	// apiParam := c.Params("*")
	//})

	// Creating a static logical code and listening to server
	app.Static("/", "./public")
	app.Listen(":" + port)
}
