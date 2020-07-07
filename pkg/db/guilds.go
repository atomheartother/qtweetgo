package db

import "context"

// GuildInfoT represents a guild's information
type GuildInfoT struct {
	ID     *string
	Prefix *string
	Lang   *string
}

// GuildInfo gets a specific guild's info from the database
func GuildInfo(id *string) (*GuildInfoT, error) {
	g := &GuildInfoT{ID: id}
	err := Pg.QueryRow(context.Background(), "SELECT lang, prefix FROM guilds WHERE id=$1 LIMIT 1", id).Scan(&g.Lang, &g.Prefix)
	if err != nil {
		return nil, err
	}
	return g, nil
}
