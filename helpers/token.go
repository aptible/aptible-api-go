package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// GetAuthURL returns the Aptible auth API URL from APTIBLE_AUTH_ROOT_URL or the default.
func GetAuthURL() string {
	if auth := os.Getenv("APTIBLE_AUTH_ROOT_URL"); auth != "" {
		return auth
	}
	return "https://auth.aptible.com"
}

// GetToken resolves an Aptible access token using the following precedence:
// 1. APTIBLE_USERNAME + APTIBLE_PASSWORD environment variables (login to auth API)
// 2. APTIBLE_ACCESS_TOKEN environment variable
// 3. ~/.aptible/tokens.json file (keyed by auth URL)
func GetToken() (string, error) {
	authURL := GetAuthURL()

	user := os.Getenv("APTIBLE_USERNAME")
	password := os.Getenv("APTIBLE_PASSWORD")
	if user != "" && password != "" {
		return loginWithUsernameAndPassword(authURL, user, password)
	}

	if token := os.Getenv("APTIBLE_ACCESS_TOKEN"); token != "" {
		return token, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to determine home directory: %w", err)
	}

	data, err := os.ReadFile(filepath.Join(home, ".aptible", "tokens.json"))
	if err != nil {
		return "", fmt.Errorf("no token found, are you logged in? Error: %w", err)
	}

	var tokens map[string]string
	if err := json.Unmarshal(data, &tokens); err != nil {
		return "", fmt.Errorf("invalid tokens.json: %w", err)
	}

	if token := tokens[authURL]; token != "" {
		return token, nil
	}

	return "", fmt.Errorf("no token found for %s, are you logged in?", authURL)
}

func loginWithUsernameAndPassword(authURL, user, password string) (string, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"expires":    43200,
		"username":   user,
		"password":   password,
		"grant_type": "password",
		"scope":      "manage",
	})
	if err != nil {
		return "", fmt.Errorf("error encoding login payload: %w", err)
	}

	req, err := http.NewRequest("POST", authURL+"/tokens", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("error creating login request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error logging in: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading login response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error decoding login response: %w", err)
	}

	token, ok := result["access_token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("no access_token in login response")
	}

	return token, nil
}
