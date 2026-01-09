package strava

import (
	"encoding/json"
	"fmt"
	"gova/internal/config"
	"io"
	"log"
	"net/http"
	url2 "net/url"
)

const (
	StravaAuthURL  = "https://www.strava.com/oauth/authorize"
	StravaTokenURL = "https://www.strava.com/oauth/token"
	StravaBaseURL  = "https://www.strava.com/api/v3"
)

type Client struct {
	httpClient *http.Client
	cfg        *config.Config
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: &http.Client{},
		cfg:        cfg,
	}
}

func (c *Client) BuildAuthURL() string {
	params := url2.Values{
		"client_id":       {c.cfg.ClientId},
		"response_type":   {"code"},
		"redirect_uri":    {c.cfg.AuthRedirectURI},
		"approval_prompt": {"force"},
		"scope":           {"activity:read_all"},
	}

	return StravaAuthURL + "?" + params.Encode()
}

func (c *Client) buildTokenURL(code string) string {
	params := url2.Values{
		"client_id":     {c.cfg.ClientId},
		"client_secret": {c.cfg.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	return StravaTokenURL + "?" + params.Encode()
}

//func (c *Client) GetCurrentAthlete() {
//	body, err := c.do("")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	sb := string(body)
//	log.Printf(sb)
//}

func (c *Client) ListActivities(before int64, after int64) ([]Activity, error) {
	resp, err := c.do(fmt.Sprintf("activities?before=%d&after=%d&per_page=10", before, after))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch activities: %w", err)
	}
	defer resp.Body.Close()

	var activities []Activity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, fmt.Errorf("failed to decode activities: %w", err)
	}

	return activities, nil
}

func (c *Client) do(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", StravaBaseURL, url), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.cfg.StravaToken))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error fetching data %s \n", err)
		return nil, err
	}

	// TODO CHECK RESPONSE CODE

	return resp, nil
}

func (c *Client) ExchangeAuth(code string) (*TokenResponse, error) {
	resp, err := http.Post(c.buildTokenURL(code), "application/json", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange auth token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to exchange auth token, status: %d, body: %s", resp.StatusCode, string(body))
	}

	var token TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, fmt.Errorf("failed to decode auth token: %w", err)
	}

	return &token, nil
}
