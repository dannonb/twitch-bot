package utils

import (
	"net/http"
	"io"
	"fmt"
	"log"
	"encoding/json"
)

type QuoteResp struct {
	Quote string `json:"quote"`
	Anime string `json:"anime"`
	Character string `json:"character"`
}

func GetQuote() *QuoteResp {
	req, _ := http.NewRequest("GET", "https://animechan.xyz/api/random", nil)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(error.Error(err))
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

		response := QuoteResp{}
		json.Unmarshal(bodyBytes, &response)

		return &response

	}

	return &QuoteResp{}

}