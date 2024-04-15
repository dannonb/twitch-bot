package main

import (
	"fmt"
	"os"

	"github.com/dannonb/twitch-bot-bumr/config"
	"github.com/dannonb/twitch-bot-bumr/utils"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/TwiN/go-away"
)

var (
	channel string
	username string
	password string
)

func init() {
	config.LoadEnv()

	channel = os.Getenv("CHANNEL")
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
}

func main() {
	client := twitch.NewClient(username, password)

	utils.GetApexStats()

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Message == "!hi" {
			client.Say(channel, "Hello, welcome to the stream!")
		}

		if goaway.IsProfane(message.Message) {
			client.Say(channel, fmt.Sprintf("@%s keep it pg!", message.User.DisplayName))
		}

		utils.MessageHandler(channel, client, message)

	})

	client.OnConnect(func() {
		fmt.Println("Connected")
	})

	client.Join(channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}