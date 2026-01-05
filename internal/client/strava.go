package client

import (
	"encoding/json"
	"fmt"
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

//func (c *Client) GetCurrentAthlete() {
//	body, err := c.do("")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	sb := string(body)
//	log.Printf(sb)
//}

func (c *Client) ListActivities(before int, after int) ([]Activity, error) {
	resp, err := c.do(fmt.Sprintf("activities?before=%d&after%d=&per_page=10", before, after))
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, url), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken))
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error fetching data %s \n", err)
		return nil, err
	}

	// TODO CHECK RESPONSE CODE

	return resp, nil
}
