package discord

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/atomheartother/qtweet/pkg/db"
	"github.com/bwmarrin/discordgo"
)

// Init initializes the discord connection
func Init() {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
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
