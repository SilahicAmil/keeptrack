package azuredevops

import (
	"changeme/internal/services/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetch PRs for ticket(s), reviewers, threads

// Fetch ALL active PR's for the user
func (c *AzureDevopsClient) FetchPullRequests(userID string) ([]models.PullRequest, error) {
	// Get all PRs for now

	// GET https://dev.azure.com/{org}/{project}/_apis/git/pullrequests
	//   ?searchCriteria.creatorId={userId}
	//   &searchCriteria.status=active
	//   &api-version=7.1

	// Get the user id. FIrst we need to store it lol.

	PRUrl := fmt.Sprintf("%s/_apis/git/pullrequests?searchCriteria.creatorId=%s&searchCriteria.status=active&api-version=7.1",
		c.BaseURL, userID)

	req, err := http.NewRequest("GET", PRUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.AuthHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.PullRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var pullRequests []models.PullRequest

	for _, item := range result.Value {

		reviewers := make([]models.Reviewers, 0, len(item.Reviewers))
		for _, rev := range item.Reviewers {
			reviewers = append(reviewers, models.Reviewers{
				DisplayName: rev.DisplayName,
				Vote:        rev.Vote,
			})
		}
		pullRequests = append(pullRequests, models.PullRequest{
			ID:          item.ID,
			Status:      item.Status,
			Title:       item.Title,
			Description: item.Description,
			Reviewers:   reviewers,
		})
	}
	return pullRequests, nil
}

func (c *AzureDevopsClient) FetchPullRequestsReviewer(userID string) ([]models.PullRequest, error) {
	// Get all PRs for now

	// GET https://dev.azure.com/{org}/{project}/_apis/git/pullrequests
	//   ?searchCriteria.creatorId={userId}
	//   &searchCriteria.status=active
	//   &api-version=7.1

	// Get the user id. FIrst we need to store it lol.

	PRUrl := fmt.Sprintf("%s/_apis/git/pullrequests?searchCriteria.reviewerId=%s&searchCriteria.status=active&api-version=7.1",
		c.BaseURL, userID)

	req, err := http.NewRequest("GET", PRUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.AuthHeader)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result models.PullRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var pullRequests []models.PullRequest

	for _, item := range result.Value {

		reviewers := make([]models.Reviewers, 0, len(item.Reviewers))
		for _, rev := range item.Reviewers {
			reviewers = append(reviewers, models.Reviewers{
				DisplayName: rev.DisplayName,
				Vote:        rev.Vote,
			})
		}
		pullRequests = append(pullRequests, models.PullRequest{
			ID:          item.ID,
			Status:      item.Status,
			Title:       item.Title,
			Description: item.Description,
			Reviewers:   reviewers,
		})
	}
	return pullRequests, nil
}
