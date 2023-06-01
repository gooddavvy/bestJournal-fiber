package envViper

import (
	"fmt"

	"github.com/spf13/viper"
)

func Config(testAllowed bool) {
	viper.AutomaticEnv()

	if testAllowed {
		// Get the value of the environment variable "REGISTER_TOKEN"
		registerToken := viper.Get("REGISTER_TOKEN")

		fmt.Println(registerToken)
	}
}
