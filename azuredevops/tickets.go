package azuredevops

import (
	"changeme/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"
)

// HTTP Client, Auth, Base Request Logic

func (c *AzureDevopsClient) queryAssignedWorkItems() ([]int, error) {

	base := fmt.Sprintf("https://dev.azure.com/%s/%s", c.cfg.Org, c.cfg.Project)
	url := base + "/_apis/wit/wiql?api-version=7.1"

	query := `{
		"query": "SELECT [System.Id], [System.Title] FROM WorkItems WHERE [System.AssignedTo] = @Me AND [System.State] <> 'Closed' ORDER BY [System.ChangedDate] DESC"
	}`

	req, err := http.NewRequest("POST", url, strings.NewReader(query))

	if err != nil {
		return nil, err
	}

	auth := c.authHeader

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result Response

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		return nil, err
	}

	var ids []int

	for _, w := range result.WorkItems {
		ids = append(ids, w.ID)

	}

	return ids, nil
}

func (c *AzureDevopsClient) queryAssignedWorkItemsData() ([]Ticket, error) {

	ids, err := c.queryAssignedWorkItems()

	idString := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")

	// Sprintf the below. use c.cfg.project and org

	baseURL := fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/wit/workitems",
		c.cfg.Org,
		c.cfg.Project,
	)

	url := fmt.Sprintf(
		"%s?ids=%s&fields=System.Id,System.Title,System.State,System.Description,System.WorkItemType&api-version=7.1",
		baseURL,
		idString,
	)

	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	auth := c.authHeader

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Basic "+auth)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result WorkItemResponse

	json.NewDecoder(resp.Body).Decode(&result)

	var tickets []Ticket

	for _, item := range result.Value {
		tickets = append(tickets, Ticket{
			ID:          item.ID,
			Title:       item.Fields.Title,
			Description: strip.StripTags(item.Fields.Description),
			State:       item.Fields.State,
		})
	}

	return tickets, nil
}

func (c *AzureDevopsClient) FetchAssignedTickets() ([]Ticket, error) {

	// Load PAT from .env
	// query azure devops
	// format into []Ticket
	// and return

	tickets, err := c.queryAssignedWorkItemsData()

	if err != nil {
		return nil, err
	}

	return tickets, nil

}

func (c *AzureDevopsClient) FetchAssignedTicketsCache() ([]Ticket, error) {

	query := `SELECT * from tickets`

	res, err := store.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer res.Close()

	var tickets []Ticket
	for res.Next() {
		var t Ticket
		err := res.Scan(&t.ID, &t.Description, &t.Title, &t.State, &t.PRLinks)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}

	err = res.Err()

	if err != nil {
		return nil, err
	}

	return tickets, nil
}
