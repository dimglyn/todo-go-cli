package main
import (
	"github.com/bwmarrin/discordgo"

	"fmt"
)

var todoRepo TodoRepo

func init() {
	todoRepo = TodoRepo{}
}



func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	query, err := parseInput(m.Content)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Tell me what to do: ")
	}

	todoRepo = handle(todoRepo, query)
	s.ChannelMessageSend(m.ChannelID, todoRepo.String())
}