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
		log.Fatalln(err)
	}

	var activities []Activity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
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

func (c *Client) ExchangeAuth(code string) TokenResponse {
	req, err := http.Post(c.buildTokenURL(code), "application/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	var resp TokenResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		fmt.Println(err)
	}

	return resp
}
