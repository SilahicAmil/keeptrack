package azuredevops

// Structs for PR, Tickets, Reviewer, Comments, Status/Column and anything else

type Config struct {
	PAT     string `json:"pat"`
	Org     string `json:"org"`
	Project string `json:"project"`
}

type Ticket struct {
	ID          int      `json:"ID"`
	Title       string   `json:"Title"`
	Description string   `json:"Description"`
	State       string   `json:"State"`
	PRLinks     []string `json:"PRLinks"`
}

type CurrentUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"emailAddress"`
}

type WorkItemRef struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Response struct {
	WorkItems []WorkItemRef `json:"workItems"`
}

type WorkItemResponse struct {
	Value []struct {
		ID     int `json:"id"`
		Fields struct {
			Title        string `json:"System.Title"`
			State        string `json:"System.State"`
			Description  string `json:"System.Description"`
			WorkItemType string `json:"System.WorkItemType"`
		} `json:"fields"`
	} `json:"value"`
}
