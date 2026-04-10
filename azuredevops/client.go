package azuredevops

import (
	"changeme/store"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AzureDevopsClient struct {
}

func (c *AzureDevopsClient) loadAndCreatePATAuth() (string, error) {

	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	azPat := os.Getenv("AZURE_PAT")

	auth := base64.StdEncoding.EncodeToString([]byte(":" + azPat))

	return auth, nil
}

func (c *AzureDevopsClient) ValidateConfig(cfg Config) error {
	// just check the orgs endpoint for validation.
	fmt.Println("Inside here all good", cfg)
	return nil
}

func (c *AzureDevopsClient) StoreConfig(cfg Config) error {

	// Get all the values from config and save it to DB
	// under config table
	query := "INSERT INTO config (id, pat, org, project)"
	_, err := store.DB.Exec(query, 1, cfg.PAT, cfg.Org, cfg.Project)
	return err
}
