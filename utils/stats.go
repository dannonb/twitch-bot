package utils

import (
	"fmt"
	"net/http"
	"os"
)



func GetApexStats() {
	key := os.Getenv("TRN_API_KEY")
	// still in test, hardcoded values for now
	req, _ := http.NewRequest("GET", "https://public-api.tracker.gg/v2/apex/standard/profile/psn/dannon", nil)
	req.Header.Set("TRN-Api-Key", key)

	client := &http.Client{}
	resp, err := client.Do(req) 

	if err != nil {
		fmt.Println(error.Error(err))
	}

	fmt.Println(resp)
	
}