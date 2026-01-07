package config

import (
	"errors"
	"os"
)

type Config struct {
	StravaToken  string
	BaseURL      string
	ClientId     string
	ClientSecret string
}

func Load() (*Config, error) {
	token := os.Getenv("STRAVA_AUTH_TOKEN")
	if token == "" {
		return nil, errors.New("STRAVA_AUTH_TOKEN environment variable not set")
	}

	baseURL := os.Getenv("STRAVA_BASE_URL")
	if baseURL == "" {
		return nil, errors.New("BASE_URL environment variable not set")
	}

	return &Config{token, baseURL, "", ""}, nil
}
