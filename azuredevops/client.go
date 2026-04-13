package azuredevops

import (
	"changeme/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AzureDevopsClient struct {
	cfg        *config.CFG
	baseURL    string
	authHeader string
}

func NewAzureDevopsClient(cfg *config.CFG) *AzureDevopsClient {

	baseURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s",
		cfg.Org,
		cfg.Project,
	)

	auth := base64.StdEncoding.EncodeToString([]byte(":" + cfg.PAT))

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

func (c *AzureDevopsClient) FetchUser() (*CurrentUser, error) {

	url := "https://app.vssps.visualstudio.com/_apis/profile/profiles/me?api-version=7.1"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	auth := c.authHeader

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth("", auth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("azure error: %s - %s", resp.Status, string(body))
	}

	var user CurrentUser
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
