package config

import (
	"errors"
	"os"
)

type Config struct {
	StravaToken     string
	ClientId        string
	ClientSecret    string
	AuthRedirectURI string
}

func Load() (*Config, error) {
	token := os.Getenv("STRAVA_AUTH_TOKEN")
	if token == "" {
		return nil, errors.New("STRAVA_AUTH_TOKEN environment variable not set")
	}

	clientId := os.Getenv("STRAVA_CLIENT_ID")
	if clientId == "" {
		return nil, errors.New("STRAVA_CLIENT_ID environment variable not set")
	}

	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	if clientSecret == "" {
		return nil, errors.New("STRAVA_CLIENT_SECRET environment variable not set")
	}

	redirectURI := os.Getenv("AUTH_REDIRECT_URI")
	if redirectURI == "" {
		return nil, errors.New("AUTH_REDIRECT_URI environment variable not set")
	}

	return &Config{token, clientId, clientSecret, redirectURI}, nil
}
