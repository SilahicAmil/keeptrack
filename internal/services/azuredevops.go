package services

import (
	"changeme/azuredevops"
	"changeme/config"
	"changeme/internal/services/models"
	"changeme/poller"
	"changeme/store"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AzureDevopsService struct {
	client      *azuredevops.AzureDevopsClient
	store       *store.SQLiteStore
	azureCFG    config.AzureCFG
	currentUser *models.CurrentUser
	isPolling   bool
	stopPoll    context.CancelFunc
	// ticketStore  *store.TicketStore
	// prStore *store.SQLiteStore
	// pollInterval time.Duration

}

// go routine poller? startPoller func
// call azuredevops.client
// then azuredevops.tickets
// then store.ticketstore?

func NewAzureDevopsService(store *store.SQLiteStore) *AzureDevopsService {
	return &AzureDevopsService{
		client: &azuredevops.AzureDevopsClient{},
		store:  store,
	}
}

func (s *AzureDevopsService) Start(ctx context.Context) {
	app := application.Get() // safe here

	//
	// TODO : UPDATE TIME BEFORE RELASE - IMPORTANT
	//
	go poller.StartTicketPoller(ctx, 30*time.Second, s.FetchAssignedTickets, func(tickets []models.Ticket) {
		fmt.Println("This fired")
		fmt.Println("tickets", tickets)
		app.Event.Emit("tickets-updated", tickets)

		// Also fetch new PR updates. Don't worry about notifs for now
		prs, _ := s.FetchPullRequests()
		app.Event.Emit("prs-update", prs)

		prsReviewer, _ := s.FetchPullRequestsReviewer()
		app.Event.Emit("prs-reviwer", prsReviewer)
	})
}

func (s *AzureDevopsService) FetchAssignedTickets() ([]models.Ticket, error) {
	return s.client.FetchAssignedTickets(s.currentUser, s.store)
}

func (s *AzureDevopsService) FetchAssignedTicketsCache() ([]models.Ticket, error) {

	return s.client.FetchAssignedTicketsCache()

}

func (s *AzureDevopsService) InitializeApp(cfg config.AzureCFG) ([]models.Ticket, error) {
	s.azureCFG = cfg

	s.client = azuredevops.NewAzureDevopsClient(&cfg)

	if err := s.client.ValidateConfig(); err != nil {
		return nil, err
	}

	if s.azureCFG.Validate {
		// Just return an empty struct back for now
		return []models.Ticket{}, nil
	}

	// Save config
	if err := s.store.StoreConfig(cfg); err != nil {
		return nil, err
	}

	user, err := s.client.FetchUser()
	if err != nil {
		return nil, err
	}

	s.currentUser = user

	if err := s.store.SaveUserToConfig(user); err != nil {
		return nil, err
	}

	tickets, err := s.client.FetchAssignedTickets(user, s.store)
	if err != nil {
		return nil, err
	}

	if err := s.store.SaveTickets(tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *AzureDevopsService) StartPolling() {
	if s.isPolling {
		// prevent dupe go routines
		return
	}

	if s.client == nil {
		log.Println("poller not started: client not initialized")
		return
	}

	// Manage the poller
	ctx, cancel := context.WithCancel(context.Background())
	s.stopPoll = cancel
	s.isPolling = true

	go func() {
		defer func() { s.isPolling = false }()

		s.Start(ctx)
	}()
}

func (s *AzureDevopsService) CheckAppState() (bool, error) {

	ready, cfg, user, err := s.store.CheckAppState()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if !ready {
		fmt.Println("not ready?")
		return false, nil
	}

	// use cfg directly (no need to rebuild)
	s.client = azuredevops.NewAzureDevopsClient(&cfg)

	// store user in memory
	s.currentUser = &user

	return true, nil
}

func (s *AzureDevopsService) OpenTicket(id int) {
	app := application.Get()

	ticketURL := fmt.Sprintf("%s/_workitems/edit/%d", s.client.BaseURL, id)
	err := app.Browser.OpenURL(ticketURL)

	if err != nil {
		app.Logger.Error("failed to open link", "error", err)
	}
}

func (s *AzureDevopsService) OpenPR(id int) {
	app := application.Get()

	ticketURL := fmt.Sprintf("%s/_git/%s/pullrequest/%d", s.client.BaseURL, s.client.CFG.Project, id)
	err := app.Browser.OpenURL(ticketURL)

	if err != nil {
		app.Logger.Error("failed to open link", "error", err)
	}
}

func (s *AzureDevopsService) FetchPullRequests() ([]models.PullRequest, error) {
	return s.client.FetchPullRequests(s.currentUser.ID)
}

func (s *AzureDevopsService) FetchPullRequestsReviewer() ([]models.PullRequest, error) {
	return s.client.FetchPullRequestsReviewer(s.currentUser.ID)
}
