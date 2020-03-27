package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token    string
	todoRepo TodoRepo
	scanner  *bufio.Scanner
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	todoRepo = TodoRepo{}
}

func main() {
	discord, err := discordgo.New("Bot " + Token)

	if err != nil {
		fmt.Println("error lol, ", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
