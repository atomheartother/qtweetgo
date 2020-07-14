package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/atomheartother/qtweet/pkg/config"
	"github.com/atomheartother/qtweet/pkg/db"
	"github.com/atomheartother/qtweet/pkg/discord"
)

func main() {
	config.InitEnv()
	err := db.Init()
	if err != nil {
		fmt.Println("Couldn't connect to the database,", err)
		return
	}
	err = discord.Init()
	if err != nil {
		fmt.Println("Couldn't connect to Discord,", err)
		return
	}
	sc := make(chan os.Signal, 1)
	fmt.Println("The bot is now running - press Ctrl-C to stop it")
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = discord.Close()
	if err != nil {
		fmt.Println("Error closing out Discord connection", err)
		err = nil
	} else {
		fmt.Println("Discord disconnection successful.")
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error closing database connection", err)
	} else {
		fmt.Println("Database closed successfully.")
	}
	return
}
