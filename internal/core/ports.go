package core

type OauthClient interface {
	BuildAuthURL() string
	ExchangeToken(string) (*TokenResponse, error)
}

type ApiClient interface {
	ListActivities(before, after int64) ([]Activity, error)
}
