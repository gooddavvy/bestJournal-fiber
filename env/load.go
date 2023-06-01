package envViper

import (
	"fmt"

	"github.com/joho/godotenv"
)

// Load will read your env file(s) and load them into ENV for this process.
//
// It also handles the edge case for errors loading, for example.
//
// Huge credits to joho for the original function, though.
func Load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}
}
