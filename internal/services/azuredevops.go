package services

import (
	"changeme/azuredevops"
	"changeme/config"
	"changeme/internal/services/models"
	"changeme/store"
	"context"
	"fmt"
)

type AzureDevopsService struct {
	client   *azuredevops.AzureDevopsClient
	store    *store.SQLiteStore
	azureCFG config.AzureCFG
	// ticketStore  *store.TicketStore
	// prStore      *store.PRStore
	// pollInterval time.Duration
	// stopPoll     context.CancelFunc
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
	// app := application.Get() // safe here

	// go poller.StartTicketPoller(ctx, 1*time.Minute, s.FetchAssignedTickets, func(tickets []models.Ticket) {
	// 	fmt.Println("This fired")
	// 	fmt.Println("tickets", tickets)
	// 	app.Event.Emit("tickets-updated", tickets)
	// })
}

func (s *AzureDevopsService) FetchAssignedTickets() ([]models.Ticket, error) {
	tickets, _ := s.client.FetchAssignedTickets()
	fmt.Println("start up ", tickets)
	return s.client.FetchAssignedTickets()
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

	if err := s.store.SaveUser(user); err != nil {
		return nil, err
	}

	tickets, err := s.client.FetchAssignedTickets()
	if err != nil {
		return nil, err
	}

	if err := s.store.SaveTickets(tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *AzureDevopsService) fetchAndUpdate() {}
