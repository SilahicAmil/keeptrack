package azuredevops

import (
	"encoding/base64"
	"os"

	"github.com/joho/godotenv"
)

type AzureDevopsClient struct{}

func (c *AzureDevopsClient) loadAndCreatePATAuth() (string, error) {

	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	azPat := os.Getenv("AZURE_PAT")

	auth := base64.StdEncoding.EncodeToString([]byte(":" + azPat))

	return auth, nil
}
