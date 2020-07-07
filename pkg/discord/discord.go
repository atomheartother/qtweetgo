package discord

import (
	"fmt"
	"log"
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
	fmt.Println("New message")
	// Ignore all messages created by any bot
	if m.Author.Bot {
		return
	}
	var guildid string = "149624121024577536"
	fmt.Println("Getting guild info...")
	info, err := db.GuildInfo(&guildid)
	fmt.Println("Got guild info")
	if err != nil {
		log.Fatal("Error getting guild info,", err)
		return
	}
	fmt.Println(info.ID, info.Prefix, info.Lang)
}
