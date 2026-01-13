package core

type OauthClient interface {
	BuildAuthURL() string
	ExchangeToken(string) (*TokenResponse, error)
	RefreshToken(string) (*TokenResponse, error)
}

type ApiClient interface {
	ListActivities(before int64, after int64) ([]Activity, error)
	GetCurrentAthlete() (*Athlete, error)
}

type TokenProvider interface {
	GetAccessToken() (string, error)
}
