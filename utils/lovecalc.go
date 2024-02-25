package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Resp struct {
	Fname      string `json:"fname"`
	Sname      string `json:"sname"`
	Percentage string `json:"percentage"`
	Result     string `json:"result"`
}

func LoveCalc(fname string, sname string) *Resp {
	key := os.Getenv("RAPID_API_KEY")
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://love-calculator.p.rapidapi.com/getPercentage?sname=%s&fname=%s", sname, fname), nil)
	req.Header.Set("X-RapidAPI-Key", key)
	req.Header.Set("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")

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

		response := Resp{}
		json.Unmarshal(bodyBytes, &response)

		return &response

	}

	return &Resp{}

}
