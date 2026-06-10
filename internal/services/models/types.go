package models

// Structs for PR, Tickets, Reviewer, Comments, Status/Column and anything else

type Ticket struct {
	ID               int    `json:"ID"`
	Title            string `json:"Title"`
	Description      string `json:"Description"`
	State            string `json:"State"`
	Tags             string `json:"Tags"`
	AssignedTo       string `json:"AssignedTo"`
	IsAssignedToMe   bool   `json:"IsAssignedToMe"`
	ChangedDate      string `json:"ChangedDate"`
	LastNotifiedDate string `json:"LastNotifiedDate"`
	// PRIds            []int  `json:"PRIds"`
}

type CurrentUser struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Email       string `json:"emailAddress"`
}

type PullRequest struct {
	ID          int         `json:"ID"`
	Status      string      `json:"Status"`
	Title       string      `json:"Title"`
	Description string      `json:"Description"`
	Reviewers   []Reviewers `json:"Reviewers"`
}

type PullRequestResponse struct {
	Value []struct {
		ID          int    `json:"pullRequestId"`
		Status      string `json:"status"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Reviewers   []struct {
			DisplayName string `json:"displayName"`
			Vote        int    `json:"vote"`
		} `json:"reviewers"`
	} `json:"value"`
}

type Reviewers struct {
	DisplayName string
	Vote        int
}

type ConnectionData struct {
	AuthenticatedUser struct {
		ID          string `json:"id"`
		DisplayName string `json:"providerDisplayName"`
		Properties  struct {
			Account struct {
				Value string `json:"$value"`
			} `json:"Account"`
		} `json:"properties"`
	} `json:"authenticatedUser"`
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
		ID     int            `json:"id"`
		Fields map[string]any `json:"fields"`
	} `json:"value"`
}
