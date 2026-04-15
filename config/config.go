package config

type AzureCFG struct {
	Provider string `json:"provider"`
	PAT      string `json:"pat"`
	Org      string `json:"org"`
	Project  string `json:"project,omitempty"`
	Validate bool   `json:"validate,omitempty"`
}

type JiraCFG struct {
	Provider string `json:"provider"`
	Email    string `json:"email"`
	APIToken string `json:"apiToken"`
	Domain   string `json:"domain"`
	Project  string `json:"project,omitempty"`
	Validate bool   `json:"validate,omitempty"`
}

// func (c *CFG) IsAzure() bool {
// 	return c.Provider == "azure"
// }

// func (c *CFG) IsJira() bool {
// 	return c.Provider == "jira"
// }
