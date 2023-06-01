package handlers

import (
	// "fmt"
	"fmt"
	"os"
	"strconv"
	"time"

	"math/rand"

	"github.com/gofiber/fiber/v2"
	envViper "github.com/gooddavvy/bestJournal-fiber/env"
	otherFs "github.com/gooddavvy/bestJournal-fiber/fs"
)

func RegisterAPI(c *fiber.Ctx) error {
	fullName := c.Query("fullName")
	PIN := c.Query("PIN")
	registerToken := envViper.Atos("REGISTER_TOKEN")
	if registerToken == "" {
		tokenNS := strconv.Itoa(rand.Intn(101))
		token := "gHi24io" + tokenNS + "JHG-iA" + tokenNS + "nnyi" + tokenNS + "admaa" + tokenNS + "bdllh" + tokenNS
		os.Setenv("REGISTER_TOKEN", token)

		// Create a response or JSON instance
		response := struct {
			Message  string `json:"message"`
			Token    string `json:"token"`
			FullName string `json:"fullName"`
			PIN      string `json:"pin"`
		}{
			Message:  "Successfully registered! Don't forget to keep the token and PIN below secret.",
			Token:    token,
			FullName: fullName,
			PIN:      PIN,
		}

		// Set the cache-control headers
		c.Set("Cache-Control", "no-store")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", time.Now().AddDate(0, 0, -1).Format(time.RFC1123))

		// Set the "Content-Type" header to "application/json"
		c.Type("application/json")

		// Send the JSON response
		return c.JSON(response)
	} else {
		return nil
	}
}

func JournalPagesAPI(c *fiber.Ctx) error {
	journalPages, err := otherFs.GetJson()

	if err != nil {
		fmt.Println("Error:", err) // Add this line for debugging
		return err
	}

	// Set the required content type
	c.Context().SetContentType("application/json")

	// Disable caching by setting appropriate cache control headers
	c.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")

	// Check if journalPages is empty
	if len(journalPages) == 0 {
		// Return empty array response
		return c.SendString("[]")
	}

	// Return the JSON response
	return c.JSON(journalPages)
}

// This is called the Main API. Don't worry, you'll see what it does.
func MainAPI(c *fiber.Ctx) error {
	apiParam := c.Params("*")

	rand.Seed(time.Now().UnixNano())

	if apiParam == "register" {
		return RegisterAPI(c)
	}
	return nil
}
