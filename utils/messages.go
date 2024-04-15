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

	if data.Command == "!commands" {
		client.Say(channel, "!lovecalc {name} | !quote | !stats")
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

	if data.Command == "!apexstats" {
		overview := GetApexStats()

		client.Say(channel, fmt.Sprintf("Kills: %s | Damage: %s | Wins: %s | Most Played Legend: %s", overview.Kills, overview.Damage, overview.Wins, overview.Legend))
	}

	if data.Command == "!age" {
		age := "30"
		client.Say(channel, fmt.Sprintf("I am %s years old.", age))
	}

	if data.Command == "!sponsor" {
		url := "https://dubby.gg"
		client.Say(channel, url)
	}

	if data.Command == "!setup" {
		monitor := "https://tinyurl.com/2s3wfp7x"
		keyboard := "https://tinyurl.com/5xn3ytvh"
		mouse := "https://tinyurl.com/bdcut9kt"
		headset := "https://tinyurl.com/kp6ce84k"
		client.Say(channel, fmt.Sprintf(`Monitor: %s | Keyboard: %s | Mouse: %s | Headset: %s`, monitor, keyboard, mouse, headset))
	}

	if data.Command == "!ranks" {
		apex := "Plat 4"
		siege := "Gold 1"
		client.Say(channel, fmt.Sprintf(`Apex: %s | Siege: %s `, apex, siege))
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