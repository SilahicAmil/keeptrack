// poller/poller.go
package poller

import (
	"context"
	"log"
	"time"
)

// Generic poller
func StartPoller[T any](
	ctx context.Context,
	interval time.Duration,
	fetch func() (T, error),
	callback func(T),
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			data, err := fetch()
			if err != nil {
				log.Println("poller error:", err)
				time.Sleep(5 * time.Second)
				continue
			}

			callback(data)
		}
	}
}
