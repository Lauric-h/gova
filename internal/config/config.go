package config

import (
	"errors"
	"os"
)

type Config struct {
	ClientId        string
	ClientSecret    string
	AuthRedirectURI string
}

func Load() (*Config, error) {
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

	return &Config{clientId, clientSecret, redirectURI}, nil
}
