package strava

import (
	"encoding/json"
	"fmt"
	"gova/internal/config"
	"gova/internal/core"
	"io"
	"net/http"
	url2 "net/url"
)

const (
	AuthURL  = "https://www.strava.com/oauth/authorize"
	TokenURL = "https://www.strava.com/oauth/token"
)

type OauthClient struct {
	cfg *config.Config
}

func NewOauthClient(cfg *config.Config) *OauthClient {
	return &OauthClient{cfg: cfg}
}

func (c *OauthClient) ExchangeToken(code string) (*core.TokenResponse, error) {
	resp, err := http.Post(c.BuildTokenURL(code), "", nil)
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

func (c *OauthClient) BuildAuthURL() string {
	params := url2.Values{
		"client_id":       {c.cfg.ClientId},
		"response_type":   {"code"},
		"redirect_uri":    {c.cfg.AuthRedirectURI},
		"approval_prompt": {"force"},
		"scope":           {"activity:read_all"},
	}

	return AuthURL + "?" + params.Encode()
}

func (c *OauthClient) BuildTokenURL(code string) string {
	params := url2.Values{
		"client_id":     {c.cfg.ClientId},
		"client_secret": {c.cfg.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	return TokenURL + "?" + params.Encode()
}
