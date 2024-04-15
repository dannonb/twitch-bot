package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Overview struct {
	Level string `json:"level"`
	Kills string `json:"kills"`
	Wins string `json:"wins"`
	Legend string `json:"legend"`
	Damage string `json:"damage"`
}

type ApexResponse struct {
	Data struct {
		Segments []StatSegment `json:"segments"`
	} `json:"data"`
}

type StatSegment struct {
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	Metadata   Metadata   `json:"metadata"`
	ExpiryData string     `json:"expiryDate"`
	Stats      Stats      `json:"stats"`
}

type Attributes struct {
	Id string `json:"id"`
}

type Metadata struct {
	Name             string `json:"name"`
	ImageUrl         string `json:"imageUrl"`
	TallImageUrl     string `json:"tallImageUrl"`
	BgImageUrl       string `json:"bgImageUrl"`
	PortraitImageUrl string `json:"portraitImageUrl"`
	LegendColor      string `json:"legendColor"`
	IsActive         bool   `json:"isActive"`
}

type Stats struct {
	Kills  Kills  `json:"kills"`
	Damage Damage `json:"damage"`
	Wins   Wins   `json:"wins"`
	Level  Level  `json:"level"`
}

type Kills struct {
	Value float32 `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type Damage struct {
	Value float32 `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type Wins struct {
	Value float32 `json:"value"`
	DisplayValue string `json:"displayValue"`
}

type Level struct {
	Value float32 `json:"value"`
	DisplayValue string `json:"displayValue"`

}

func GetApexStats() *Overview {
	key := os.Getenv("TRN_API_KEY")
	// still in test, hardcoded values for now
	req, _ := http.NewRequest("GET", "https://public-api.tracker.gg/v2/apex/standard/profile/psn/ATObumr", nil)
	req.Header.Set("TRN-Api-Key", key)

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

		var response ApexResponse

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			log.Fatal(err)
		}

		var (
			legend string
			kills float32
		)

		kills = 0
		overview := &Overview{}

		for _, segment := range response.Data.Segments {
			if segment.Type == "overview" {
				overview.Kills = segment.Stats.Kills.DisplayValue
				overview.Level = segment.Stats.Level.DisplayValue
				overview.Wins = segment.Stats.Wins.DisplayValue
				overview.Damage = segment.Stats.Damage.DisplayValue
			}
			if segment.Type == "legend" && segment.Stats.Kills.Value > kills {
				kills = segment.Stats.Kills.Value
				legend = segment.Metadata.Name
			}
		}

		overview.Legend = legend

		return overview

	}

	return &Overview{}

}
