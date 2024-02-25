package utils

import (
	"fmt"
	"strings"

	
	"github.com/gempir/go-twitch-irc/v4"
)

type MsgData struct {
	Command string
	Args []string
	Username string
	UserAt string
}

func MessageHandler(channel string, client *twitch.Client, message twitch.PrivateMessage) {
	data := CommandParser(message)
	fmt.Println(data)
	if data.Command == "!stats" {
		// stats := utils.GetApexStats()
		client.Say(channel, data.UserAt + " No stats available at the moment")
	}

	if data.Command == "!lovecalc" {
		if len(data.Args) < 1 {
			client.Say(channel, "Please provide the name of your match.")
			return 
		}
		first := data.Username
		second := data.Args[0]

		resp := LoveCalc(first, second)

		client.Say(channel, fmt.Sprintf("%s: %s%% match with %s. %s\n", resp.Fname, resp.Percentage, resp.Sname, resp.Result))

	}

	if data.Command == "!quote" {
		resp := GetQuote()

		client.Say(channel, fmt.Sprintf("%q' - %s, %s", resp.Quote, resp.Character, resp.Anime))
	}
}

func CommandParser(message twitch.PrivateMessage) MsgData {
	msg := strings.Split(message.Message, " ")
	data := MsgData{}
	if len(msg) > 1 {
		data.Command = msg[0]
		data.Args = msg[1:]
		data.Username = message.User.DisplayName
		data.UserAt = "@" + message.User.DisplayName
	} else {
		data.Command = msg[0]
		data.Args = make([]string, 0)
		data.Username = message.User.DisplayName
		data.UserAt = "@" + message.User.DisplayName
	}
	return data
}