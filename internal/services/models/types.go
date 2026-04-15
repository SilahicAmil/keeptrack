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
		ID     int            `json:"id"`
		Fields map[string]any `json:"fields"`
	} `json:"value"`
}
