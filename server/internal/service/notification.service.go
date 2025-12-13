package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/dto"
	"github.com/jirugutema/kaishare/internal/repository"
	"github.com/jirugutema/kaishare/pkg"
)

func CreateNotificationService(notification dto.CreateNotificationDTO) (uuid.UUID, error) {
	notification.CreatedAt = time.Now()
	notificationID, err := pkg.IDGenerator()
	if err != nil {
		return uuid.Nil, err
	}
	notification.ID = notificationID

	if notification.UserID != uuid.Nil {
		exists, err := repository.NotificationUserExists(notification.UserID)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, ErrUserNotExist
		}
	}

	if err := repository.CreateNotification(notification); err != nil {
		return uuid.Nil, err
	}

	return notification.ID, nil
}
