package services

import (
	"changeme/azuredevops"
	"changeme/config"
	"changeme/poller"
	"changeme/store"
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AzureDevopsService struct {
	client *azuredevops.AzureDevopsClient
	store  *store.SQLiteStore
	cfg    config.CFG
	// ticketStore  *store.TicketStore
	// prStore      *store.PRStore
	// pollInterval time.Duration
	// stopPoll     context.CancelFunc
}

// go routine poller? startPoller func
// call azuredevops.client
// then azuredevops.tickets
// then store.ticketstore?

func NewAzureDevopsService() *AzureDevopsService {
	return &AzureDevopsService{
		client: &azuredevops.AzureDevopsClient{},
	}
}

func (s *AzureDevopsService) Start(ctx context.Context) {
	app := application.Get() // safe here

	go poller.StartTicketPoller(ctx, 1*time.Minute, s.FetchAssignedTickets, func(tickets []azuredevops.Ticket) {
		fmt.Println("This fired")
		fmt.Println("tickets", tickets)
		app.Event.Emit("tickets-updated", tickets)
	})
}

func (s *AzureDevopsService) FetchAssignedTickets() ([]azuredevops.Ticket, error) {
	tickets, _ := s.client.FetchAssignedTickets()
	fmt.Println("start up ", tickets)
	return s.client.FetchAssignedTickets()
}

func (s *AzureDevopsService) FetchAssignedTicketsCache() ([]azuredevops.Ticket, error) {

	return s.client.FetchAssignedTicketsCache()

}

func (s *AzureDevopsService) InitializeApp(cfg config.CFG) ([]azuredevops.Ticket, error) {
	s.cfg = cfg

	s.client = azuredevops.NewAzureDevopsClient(&cfg)

	if err := s.client.ValidateConfig(); err != nil {
		return nil, err
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

	// if err := s.store.SaveTickets(tickets); err != nil {
	// 	return nil, err
	// }

	return tickets, nil
}

func (s *AzureDevopsService) fetchAndUpdate() {}
