package service

import (
	"encoding/json"
	"fmt"
	"gova/internal/strava"
	"os"
	"path/filepath"
	"time"
)

type AuthService struct {
	Client *strava.Client
}

type Credentials struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func NewAuthService(client *strava.Client) *AuthService {
	return &AuthService{client}
}

func (s *AuthService) GetTokenFromCode(code string) error {
	tokenResponse := s.Client.ExchangeAuth(code)

	if err := s.storeToken(tokenResponse.AccessToken, tokenResponse.RefreshToken, tokenResponse.ExpiresAt); err != nil {
		return fmt.Errorf("failed to store token: %w", err)
	}

	return nil
}

func (s *AuthService) storeToken(accessToken string, refreshToken string, expiresAt int64) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "gova")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	credPath := filepath.Join(configDir, "credentials.json")
	file, err := os.OpenFile(credPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to create credentials file: %w", err)
	}
	defer file.Close()

	credentials := Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Unix(expiresAt, 0),
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(credentials); err != nil {
		return fmt.Errorf("failed to encode credentials: %w", err)
	}

	return nil
}
