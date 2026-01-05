package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	authToken  string
}

func NewClient(baseURL string, authToken string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
		authToken:  authToken,
	}
}

func (c *Client) GetCurrentAthlete() {
	body, err := c.do("")
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func (c *Client) ListActivities(before int, after int) {
	body, err := c.do(fmt.Sprintf("activities?before=%d&after%d=&per_page=10", before, after))
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func (c *Client) do(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, url), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("client: error making http request: %s\n", err)
	}

	return io.ReadAll(resp.Body)
}
