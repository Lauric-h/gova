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
	httpClient    *http.Client
	cfg           *config.Config
	tokenProvider core.TokenProvider
}

func NewClient(cfg *config.Config, t core.TokenProvider) *Client {
	return &Client{
		httpClient:    &http.Client{},
		cfg:           cfg,
		tokenProvider: t,
	}
}

func (c *Client) GetCurrentAthlete() (*core.Athlete, error) {
	resp, err := c.do("/athlete")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch athlete: %w", err)
	}
	defer resp.Body.Close()

	var athlete core.Athlete
	if err := json.NewDecoder(resp.Body).Decode(&athlete); err != nil {
		return nil, fmt.Errorf("failed to decode athlete: %w", err)
	}

	return &athlete, nil
}

func (c *Client) ListActivities(before int64, after int64) ([]core.Activity, error) {
	resp, err := c.do(fmt.Sprintf("activities?before=%d&after=%d&per_page=100", before, after))
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
	token, err := c.tokenProvider.GetAccessToken()
	if err != nil || token == "" {
		return nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", BaseURL, url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
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
