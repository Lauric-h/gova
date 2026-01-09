package strava

import (
	"encoding/json"
	"fmt"
	"gova/internal/config"
	"gova/internal/core"
	"io"
	"net/http"
)

const (
	BaseURL = "https://www.strava.com/api/v3"
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

//func (c *Client) GetCurrentAthlete() {
//	body, err := c.do("")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	sb := string(body)
//	log.Printf(sb)
//}

func (c *Client) ListActivities(before int64, after int64) ([]core.Activity, error) {
	resp, err := c.do(fmt.Sprintf("activities?before=%d&after=%d&per_page=10", before, after))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch activities: %w", err)
	}
	defer resp.Body.Close()

	var activities []core.Activity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, fmt.Errorf("failed to decode activities: %w", err)
	}

	return activities, nil
}

func (c *Client) do(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", BaseURL, url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.cfg.StravaToken))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf("failed to execute request: %d %s", resp.StatusCode, string(body))
	}

	return resp, nil
}
