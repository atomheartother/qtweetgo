package config

import (
	"fmt"
	"os"
)

// Prefix holds the default prefix for the bot
var Prefix string

// InitEnv gets the env values for configuration
func InitEnv() {
	p, exists := os.LookupEnv("PREFIX")
	if !exists {
		fmt.Println("Config Warning: PREFIX variable not set")
	}
	Prefix = p
}
