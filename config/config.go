package config

type CFG struct {
	Provider string `json:"provider,omitempty"`
	// Azure
	PAT string `json:"pat,omitempty"`
	Org string `json:"org,omitempty"`

	// Jira (Later)
	Email    string `json:"email,omitempty"`
	APIToken string `json:"apiToken,omitempty"`
	Domain   string `json:"domain,omitempty"`

	// Common
	Project string `json:"project,omitempty"`
}

func (c *CFG) IsAzure() bool {
	return c.Provider == "azure"
}

func (c *CFG) IsJira() bool {
	return c.Provider == "jira"
}
