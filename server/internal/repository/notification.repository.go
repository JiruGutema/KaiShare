// Package repository interacts with database
package repository

import (
	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/config"
	"github.com/jirugutema/kaishare/internal/dto"
)

func CreateNotification(notification dto.CreateNotificationDTO) error {
	query := `
        INSERT INTO notifications (
            id,relation_link, title, content, read, created_at, user_id
        ) VALUES ($1,$2,$3,$4,$5,$6,$7)
    `
	_, err := config.DB.Exec(
		query,

		notification.ID,
		notification.RelationLink,
		notification.Title,
		notification.Content,
		notification.Read,
		notification.CreatedAt,
		notification.UserID,
	)
	return err
}

func NotificationUserExists(userID uuid.UUID) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", userID).Scan(&exists)
	return exists, err
}

func DeleteNotification(notificationID uuid.UUID) (bool, error){
	query := "DELETE FROM notifications WHERE id=$1"
	res, err := config.DB.Exec(query, notificationID)

	return res, err
}
