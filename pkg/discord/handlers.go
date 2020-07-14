package discord

import (
	"strings"

	"github.com/atomheartother/qtweet/pkg/commands"
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
	info, _ := db.GuildInfo(&m.GuildID)
	// if err != nil {
	// 	fmt.Println("Error getting guild info,", err)
	// 	return
	// }
	var prefix *string
	if info != nil && info.Prefix != nil {
		prefix = info.Prefix
	} else {
		prefix = &config.Prefix
	}
	if !strings.HasPrefix(m.Content, *prefix) {
		return
	}
	commands.Parse(strings.TrimPrefix(m.Content, *prefix))
}
