package core

type OauthClient interface {
	BuildAuthURL() string
	BuildTokenURL(code string) string
	ExchangeToken(string) (*TokenResponse, error)
}

type ApiClient interface {
	ListActivities(before, after int64) ([]Activity, error)
}
