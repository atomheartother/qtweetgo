package discord

import (
	"fmt"
	"strings"

	"github.com/atomheartother/qtweet/pkg/config"
	"github.com/atomheartother/qtweet/pkg/db"
	"github.com/bwmarrin/discordgo"
)

// Called when a new message is created
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
	prefix := info.Prefix
	if prefix == nil {
		prefix = &config.Prefix
	}
	if !strings.HasPrefix(m.Content, *prefix) {
		return
	}

}
