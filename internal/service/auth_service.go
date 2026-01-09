package service

import (
	"gova/internal/strava"
	"time"
)

type AuthService struct {
	Client *strava.Client
}

func NewAuthService(client *strava.Client) *AuthService {
	return &AuthService{client}
}

func (s *AuthService) GetTokenFromCode(code string) {
	tokenResponse := s.Client.ExchangeAuth(code)

	s.storeToken(tokenResponse.AccessToken, tokenResponse.RefreshToken, tokenResponse.ExpiresAt)

	// -> -> -> No right scope -> abort
	// -> -> -> Store token + refresh token + exp date in config file

}

func (s *AuthService) storeToken(accessToken string, refreshToken string, expiresAt time.Time) {
	// write to credentials.json
}
