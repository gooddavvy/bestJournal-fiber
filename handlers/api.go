package handlers

import "github.com/gofiber/fiber/v2"

func JournalPagesAPI(c *fiber.Ctx) error {
	apiParam := c.Params("*")
	if apiParam == "journalPages" {
	}
	return nil
}

func MainAPI(c *fiber.Ctx) error {
	apiParam := c.Params("*")
	if apiParam == "journalPages" {
	}
	return nil
}
