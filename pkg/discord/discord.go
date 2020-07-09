package discord

import (
	"fmt"
	"os"

	"github.com/atomheartother/qtweet/pkg/db"
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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by any bot
	if m.Author.Bot {
		return
	}
	info, err := db.GuildInfo(&m.GuildID)
	if err != nil {
		fmt.Println("Error getting guild info,", err)
		return
	}
	fmt.Println(info.ID, info.Prefix, info.Lang)
}
