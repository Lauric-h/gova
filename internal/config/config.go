package config

import (
	"errors"
	url2 "net/url"
	"os"
)

type Config struct {
	StravaToken     string
	ClientId        string
	AuthRedirectURI string
}

const (
	StravaAuthURL  = "https://www.strava.com/oauth/authorize"
	StravaTokenURL = "https://www.strava.com/oauth/token"
	StravaBaseURL  = "https://www.strava.com/api/v3"
)

func Load() (*Config, error) {
	token := os.Getenv("STRAVA_AUTH_TOKEN")
	if token == "" {
		return nil, errors.New("STRAVA_AUTH_TOKEN environment variable not set")
	}

	clientId := os.Getenv("STRAVA_CLIENT_ID")
	if clientId == "" {
		return nil, errors.New("STRAVA_CLIENT_ID environment variable not set")
	}

	redirectURI := os.Getenv("AUTH_REDIRECT_URI")
	if redirectURI == "" {
		return nil, errors.New("AUTH_REDIRECT_URI environment variable not set")
	}

	return &Config{token, clientId, redirectURI}, nil
}

func (c *Config) BuildAuthURL() string {
	params := url2.Values{
		"client_id":       {c.ClientId},
		"response_type":   {"code"},
		"redirect_uri":    {c.AuthRedirectURI},
		"approval_prompt": {"force"},
		"scope":           {"activity:read_all"},
	}

	return StravaAuthURL + "?" + params.Encode()
}
