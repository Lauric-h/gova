package strava

import (
	"encoding/json"
	"fmt"
	"gova/internal/config"
	"gova/internal/core"
	"io"
	"net/http"
)

type OauthClient struct {
	cfg *config.Config
}

func (c *OauthClient) BuildAuthURL() string {
	//TODO implement me
	panic("implement me")
}

func NewOauthClient(cfg *config.Config) *OauthClient {
	return &OauthClient{cfg: cfg}
}

func (c *OauthClient) ExchangeToken(code string) (*core.TokenResponse, error) {
	resp, err := http.Post(c.buildTokenURL(code), "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange auth token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to exchange auth token, status: %d, body: %s", resp.StatusCode, string(body))
	}

	var token core.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to decode auth token: %w", err)
	}

	return &token, nil
}
