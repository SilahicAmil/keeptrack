// poller/poller.go
package poller

import (
	"changeme/internal/services/models"
	"context"
	"log"
	"time"
)

func StartTicketPoller(ctx context.Context, interval time.Duration, fetch func() ([]models.Ticket, error), callback func([]models.Ticket)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				tickets, err := fetch()
				if err == nil {
					callback(tickets)
				} else {
					log.Fatal("error fetching tickets")
				}
			}
		}
	}()
}
