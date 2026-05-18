package services

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/services/notifications"
)

type NotificationService struct {
	*notifications.NotificationService
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		NotificationService: notifications.New(),
	}
}

func (n *NotificationService) SystemNotification(id, title, body, subtitle string) (bool, error) {
	authorized, err := n.CheckNotificationAuthorization()
	if err != nil {
		return false, err
	}

	if !authorized {
		authorized, err = n.RequestNotificationAuthorization()
		if err != nil {
			return false, fmt.Errorf("request notification authorization failed: %w", err)
		}
		if !authorized {
			return false, nil
		}
	}

	err = n.SendNotification(notifications.NotificationOptions{
		ID:       id,
		Title:    title,
		Body:     body,
		Subtitle: subtitle,
	})
	if err != nil {
		return false, fmt.Errorf("send notification failed: %w", err)
	}

	return true, nil
}
