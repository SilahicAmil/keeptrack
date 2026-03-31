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

func (n *NotificationService) SystemNotification(id, title, body, subtitle string) {
	n.SendNotification(notifications.NotificationOptions{
		ID:       id,
		Title:    title,
		Body:     body,
		Subtitle: subtitle,
	})
}
