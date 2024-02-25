package config

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var oauth2Config *clientcredentials.Config

func GetAccessToken(client string, secret string) string {


	oauth2Config = &clientcredentials.Config{
		ClientID:     client,
		ClientSecret: secret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Access token: %s\n", token.AccessToken)

	return token.AccessToken
}