package envViper

import "github.com/spf13/viper"

func Atos(key string) string {
	importedKey := viper.Get(key)
	strVal, _ := importedKey.(string)
	return strVal
}
