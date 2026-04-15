package azuredevops

import (
	"changeme/internal/services/models"
	"changeme/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// HTTP Client, Auth, Base Request Logic

func (c *AzureDevopsClient) queryAssignedWorkItems() ([]int, error) {

	base := fmt.Sprintf(
		"https://dev.azure.com/%s/%s",
		c.cfg.Org,
		c.cfg.Project,
	)

	url := base + "/_apis/wit/wiql?api-version=7.1"

	query := `{
		"query": "SELECT [System.Id] FROM WorkItems WHERE [System.AssignedTo] = @Me AND [System.State] <> 'Closed'"
	}`

	req, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.authHeader)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var ids []int
	for _, w := range result.WorkItems {
		ids = append(ids, w.ID)
	}

	return ids, nil
}

func getField(fields map[string]any, key string) string {
	v, ok := fields[key]
	if !ok || v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	default:
		return fmt.Sprintf("%v", val)
	}
}

func getAssignedTo(fields map[string]any) string {
	v, ok := fields["System.AssignedTo"]
	if !ok || v == nil {
		return ""
	}

	m, ok := v.(map[string]any)
	if !ok {
		return fmt.Sprintf("%v", v)
	}

	if name, ok := m["displayName"]; ok {
		return fmt.Sprintf("%v", name)
	}

	return ""
}

func (c *AzureDevopsClient) queryAssignedWorkItemsData(user *models.CurrentUser) ([]models.Ticket, error) {

	ids, err := c.queryAssignedWorkItems()
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return []models.Ticket{}, nil
	}

	idString := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")

	baseURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/wit/workitems",
		c.cfg.Org,
		c.cfg.Project,
	)

	fields := strings.Join([]string{
		"System.Id",
		"System.Title",
		"System.State",
		"System.Description",
		"System.WorkItemType",
		"System.Tags",
		"System.AssignedTo",
		"System.ChangedDate",
	}, ",")

	url := fmt.Sprintf(
		"%s?ids=%s&fields=%s&api-version=7.1",
		baseURL,
		idString,
		fields,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.authHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.WorkItemResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var tickets []models.Ticket

	for _, item := range result.Value {

		assignedTo := getAssignedTo(item.Fields)
		fmt.Println(strip.StripTags(getField(item.Fields, "System.Description")))
		tickets = append(tickets, models.Ticket{
			ID:             item.ID,
			Title:          getField(item.Fields, "System.Title"),
			State:          getField(item.Fields, "System.State"),
			Description:    strip.StripTags(getField(item.Fields, "System.Description")),
			Tags:           getField(item.Fields, "System.Tags"),
			AssignedTo:     assignedTo,
			IsAssignedToMe: assignedTo == user.DisplayName,
			ChangedDate:    getField(item.Fields, "System.ChangedDate"),
		})
	}

	return tickets, nil
}

func (c *AzureDevopsClient) FetchAssignedTickets(user *models.CurrentUser) ([]models.Ticket, error) {

	// Load PAT from .env
	// query azure devops
	// format into []Ticket
	// and return

	tickets, err := c.queryAssignedWorkItemsData(user)

	if err != nil {
		return nil, err
	}

	return tickets, nil

}

func (c *AzureDevopsClient) FetchAssignedTicketsCache() ([]models.Ticket, error) {

	query := `
	SELECT
		id,
		title,
		IFNULL(description, ''),
		state,
		tags,
		assigned_to,
		is_assigned_to_me,
		changed_date,
		last_notified_date
	FROM tickets
	`

	res, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var tickets []models.Ticket

	for res.Next() {
		var t models.Ticket

		err := res.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.State,
			&t.Tags,
			&t.AssignedTo,
			&t.IsAssignedToMe,
			&t.ChangedDate,
			&t.LastNotifiedDate,
		)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, t)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}
