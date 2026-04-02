package services

import (
	"changeme/azuredevops"
	"changeme/poller"
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AzureDevopsService struct {
	client *azuredevops.AzureDevopsClient
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

func (s *AzureDevopsService) fetchAndUpdate() {}
