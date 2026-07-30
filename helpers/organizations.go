package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetOrganizationFromAuthAPI fetches the user's organization ID from the Aptible auth API.
// It asserts exactly one organization exists for the authenticated user.
func GetOrganizationFromAuthAPI(token string, authURL string) (string, error) {
	if authURL == "" {
		authURL = "https://auth.aptible.com"
	}

	req, err := http.NewRequest("GET", authURL+"/organizations", nil)
	if err != nil {
		return "", fmt.Errorf("error creating organization request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching organizations: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error fetching organizations: received status %d", resp.StatusCode)
	}

	var result struct {
		Embedded struct {
			Organizations []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"organizations"`
		} `json:"_embedded"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding organizations response: %w", err)
	}

	orgs := result.Embedded.Organizations
	if len(orgs) == 0 {
		return "", fmt.Errorf("no organizations found")
	}
	if len(orgs) > 1 {
		names := make([]string, len(orgs))
		for i, org := range orgs {
			names[i] = fmt.Sprintf("%s (org_id: %s)", org.Name, org.ID)
		}
		return "", fmt.Errorf("multiple organizations found, unable to determine a default: %s", strings.Join(names, ", "))
	}
	return orgs[0].ID, nil
}
