package config

import (
	"fmt"
	"os"
	"regexp"
)

// Prefix holds the default prefix for the bot
var Prefix string

// CommandRegexp simply captures a command and the rest of the text
var CommandRegexp *regexp.Regexp = regexp.MustCompile("^\\s*(\\S+)\\s*(.*)$")

// CommandArgsRegexp is the regex we use to parse our command's arguments and options. We compile it at init time
var CommandArgsRegexp *regexp.Regexp = regexp.MustCompile("--(\\w+)(=\"(.*?)\"|=(\\S+))?|\"(.*?)\"|(\\S+)")

// InitEnv gets the env values for configuration
func InitEnv() {
	p, exists := os.LookupEnv("PREFIX")
	if !exists {
		fmt.Println("Config Warning: PREFIX variable not set, using '!!' fallback")
		Prefix = "!!"
	} else {
		Prefix = p
	}
}
