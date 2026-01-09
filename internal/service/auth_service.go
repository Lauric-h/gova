package service

import (
	"encoding/json"
	"fmt"
	"gova/internal/core"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type AuthService struct {
	oauthClient core.OauthClient
}

type Credentials struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type OAuthResult struct {
	Code  string
	Error error
}

func NewAuthService(client core.OauthClient) *AuthService {
	return &AuthService{client}
}

func (s *AuthService) GetAccessToken() (string, error) {
	token, err := s.GetCredentials()
	if err != nil || token.AccessToken == "" {
		return "", err
	}

	return token.AccessToken, nil
}

func (s *AuthService) GetCredentials() (*Credentials, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	filePath := filepath.Join(homeDir, ".config", "gova", "credentials.json")
	j, err := os.Open(filePath)
	if j == nil || err != nil {
		return nil, fmt.Errorf("failed to find credentials file")
	}
	defer j.Close()

	var credentials Credentials
	if err := json.NewDecoder(j).Decode(&credentials); err != nil {
		return nil, fmt.Errorf("failed to parse credentials file")
	}

	if credentials.AccessToken == "" || credentials.RefreshToken == "" {
		return nil, fmt.Errorf("invalid credentials file")
	}

	if credentials.ExpiresAt.Before(time.Now()) {
		// Refresh Token
	}

	return &credentials, nil
}

func (s *AuthService) GetTokenFromCode(code string) error {
	tokenResponse, err := s.oauthClient.ExchangeToken(code)
	if err != nil {
		return fmt.Errorf("failed to exchange auth token, code: %s, error: %s", code, err.Error())
	}

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

func (s *AuthService) BuildLoginUrl() string {
	return s.oauthClient.BuildAuthURL()
}

func (s *AuthService) StartOAuthFlow() (*OAuthResult, error) {
	resultChan := make(chan OAuthResult, 1)

	mux := http.NewServeMux()
	mux.HandleFunc("/exchange_token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "code is required", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Authentication successful, you can close this window."))

		resultChan <- OAuthResult{Code: code, Error: nil}
	})

	server := &http.Server{Addr: ":8085", Handler: mux}

	err := exec.Command("open", s.BuildLoginUrl()).Start()
	if err != nil {
		fmt.Println("Could not open browser", err.Error())
	}

	go func() {
		_ = server.ListenAndServe()
	}()
	defer server.Close()

	select {
	case result := <-resultChan:
		return &result, nil
	case <-time.After(time.Minute * 3):
		return nil, fmt.Errorf("oauth timeout")
	}
}
