package azuredevops

import (
	"changeme/config"
	"changeme/internal/services/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type AzureDevopsClient struct {
	cfg        *config.AzureCFG
	baseURL    string
	authHeader string
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
		cfg:        cfg,
		baseURL:    baseURL,
		authHeader: auth,
	}
}

func (c *AzureDevopsClient) ValidateConfig() error {
	// just check the orgs endpoint for validation.
	fmt.Println("Inside here all good", c.cfg)
	return nil
}

func (c *AzureDevopsClient) FetchUser() (*models.CurrentUser, error) {

	url := fmt.Sprintf(
		"https://dev.azure.com/%s/_apis/projects?api-version=7.1",
		c.cfg.Org,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := c.authHeader

	log.Println("AUTH HEADER FULL ", c.authHeader)
	log.Println("PAT LEN:", len(c.cfg.PAT))

	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("azure error: %s - %s", resp.Status, string(body))
	}

	var user models.CurrentUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
