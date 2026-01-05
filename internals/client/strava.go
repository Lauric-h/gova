package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetCurrentAthlete() {
	authToken := os.Getenv("STRAVA_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("STRAVA_AUTH_TOKEN environment variable not set")
	}

	req, err := http.NewRequest(http.MethodGet, "https://www.strava.com/api/v3/athlete", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func GetActivityList() {
	authToken := os.Getenv("STRAVA_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("STRAVA_AUTH_TOKEN environment variable not set")
	}

	req, err := http.NewRequest(http.MethodGet, "https://www.strava.com/api/v3/athlete/activities?before=1767552000&after=1766956800&per_page=10", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
