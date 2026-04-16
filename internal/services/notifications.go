package services

import (
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
			return false, err
		}

		// still not authorized → stop here
		if !authorized {
			return false, nil
		}
	}

	n.SendNotification(notifications.NotificationOptions{
		ID:       id,
		Title:    title,
		Body:     body,
		Subtitle: subtitle,
	})

	return true, nil
}
