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
	client         *azuredevops.AzureDevopsClient
	store          *store.SQLiteStore
	azureCFG       config.AzureCFG
	currentUser    *models.CurrentUser
	isPolling      bool
	stopPoll       context.CancelFunc
	prCommentCache map[int]int
	// ticketStore  *store.TicketStore
	// prStore      *store.PRStore
	// pollInterval time.Duration

}

// go routine poller? startPoller func
// call azuredevops.client
// then azuredevops.tickets
// then store.ticketstore?

func NewAzureDevopsService(store *store.SQLiteStore) *AzureDevopsService {
	return &AzureDevopsService{
		client:         &azuredevops.AzureDevopsClient{},
		store:          store,
		prCommentCache: make(map[int]int),
	}
}

func (s *AzureDevopsService) Start(ctx context.Context) {
	app := application.Get() // safe here

	//
	// TODO : UPDATE TIME BEFORE RELASE - IMPORTANT. Future make this an app settings
	//
	go poller.StartPoller(ctx, 30*time.Second, s.FetchAssignedTickets, func(tickets []models.Ticket) {
		fmt.Println("This fired")
		fmt.Println("tickets", tickets)
		app.Event.Emit("tickets-updated", tickets)

	})

	// go poller.StartPoller(ctx, 60*time.Second, s.FetchActivePRs, func(prs []models.PullRequest) {
	// 	app.Event.Emit("pr-updated", prs)
	// })
}

// ------ PULL REQUESTS ------

func (s *AzureDevopsService) FetchActivePRs() ([]models.PullRequest, error) {
	return s.client.FetchActivePRs(s.currentUser)
}

func (s *AzureDevopsService) filterChangedPRs(
	prs []models.PullRequest,
) []models.PullRequest {

	var changed []models.PullRequest

	for _, pr := range prs {
		prev, exists := s.prCommentCache[pr.ID]

		if !exists || pr.CommentCount != prev {
			s.prCommentCache[pr.ID] = pr.CommentCount
			changed = append(changed, pr)
		}
	}

	return changed
}

// ------ TICKETS ------

func (s *AzureDevopsService) FetchAssignedTickets() ([]models.Ticket, error) {
	return s.client.FetchAssignedTickets(s.currentUser, s.store)
}

func (s *AzureDevopsService) FetchAssignedTicketsCache() ([]models.Ticket, error) {

	return s.client.FetchAssignedTicketsCache()

}

// ------ APP STATE ------

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
