package azuredevops

import (
	"changeme/config"
	"changeme/internal/services/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type AzureDevopsClient struct {
	CFG        *config.AzureCFG
	BaseURL    string
	AuthHeader string
}

func NewAzureDevopsClient(cfg *config.AzureCFG) *AzureDevopsClient {

	baseURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s",
		cfg.Org,
		cfg.Project,
	)

	rawPAT := strings.TrimSpace(cfg.PAT)

	auth := "Basic " + base64.StdEncoding.EncodeToString(
		[]byte(":"+rawPAT),
	)

	return &AzureDevopsClient{
		CFG:        cfg,
		BaseURL:    baseURL,
		AuthHeader: auth,
	}
}

func (c *AzureDevopsClient) ValidateConfig() error {
	// just check the orgs endpoint for validation.
	fmt.Println("Inside here all good", c.CFG)
	return nil
}

func (c *AzureDevopsClient) FetchUser() (*models.CurrentUser, error) {

	url := fmt.Sprintf(
		"https://dev.azure.com/%s/_apis/connectionData?api-version=7.1-preview.1",
		c.CFG.Org,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.AuthHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("azure error: %s - %s", resp.Status, string(body))
	}

	var conn models.ConnectionData
	if err := json.Unmarshal(body, &conn); err != nil {
		return nil, err
	}

	return &models.CurrentUser{
		ID:          conn.AuthenticatedUser.ID,
		DisplayName: conn.AuthenticatedUser.DisplayName,
		Email:       conn.AuthenticatedUser.Properties.Account.Value,
	}, nil
}
