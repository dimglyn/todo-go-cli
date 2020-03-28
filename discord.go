package main
import (
	"github.com/bwmarrin/discordgo"

	"fmt"
)

var todoRepo TodoRepo

func init() {
	todoRepo = TodoRepo{}
}

type DiscordPayload struct { 
		session *discordgo.Session
		message *discordgo.MessageCreate
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "quit" || m.Content == "exit" {
		fmt.Println("Bye bye")
		return
	}

	query, err := parseInput(m.Content)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Tell me what to do: ")
	}

	todoRepo, _ = handle(todoRepo, query, &DiscordPayload{
		session: s,
		message: m,
	})
}


func (dp *DiscordPayload) respond(res string) {
	fmt.Println(res)
	dp.session.ChannelMessageSend(dp.message.ChannelID, res)
}