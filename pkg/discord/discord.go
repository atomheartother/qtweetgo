package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

// Dg is the global Discord session
var Dg *discordgo.Session

// Init initializes the discord connection
func Init() error {
	Dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return err
	}

	Dg.AddHandler(messageCreate)

	return Dg.Open()
}

// Close closes out the discord client connection
func Close() error {
	return Dg.Close()
}
