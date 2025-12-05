package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jirugutema/gopastebin/internal/config"
	"github.com/jirugutema/gopastebin/internal/dto"
	"github.com/jirugutema/gopastebin/pkg"
)

var ErrUserNotExist = errors.New("user does not exist")

func CreatePasteService(paste dto.PasteDTO) (uuid.UUID, error) {
	paste.CreatedAt = time.Now()

	// Generate paste ID
	pasteID, err := pkg.IDGenerator()
	if err != nil {
		return uuid.Nil, err
	}
	paste.ID = pasteID

	if paste.UserID != nil {
		var exists bool
		err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", *paste.UserID).Scan(&exists)
		if err != nil {
			return uuid.Nil, err
		}
		if !exists {
			return uuid.Nil, ErrUserNotExist
		}
	}

	query := `
        INSERT INTO pastes (
            id, title, content, language, password, burn_after_read, expires_at, created_at, views, user_id, is_public
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
    `
	_, err = config.DB.Exec(
		query,
		paste.ID,
		paste.Title,
		paste.Content,
		paste.Language,
		paste.Password,
		paste.BurnAfterRead,
		paste.ExpiresAt,
		paste.CreatedAt,
		paste.Views,
		paste.UserID,
		paste.IsPublic,
	)
	if err != nil {
		return uuid.Nil, err
	}

	return paste.ID, nil
}

func GetPasteService(pasteID uuid.UUID) (dto.PasteDTO, error) {
	var res dto.PasteDTO

	query := `
        SELECT id, title, content, language, password, burn_after_read, expires_at, created_at, views, user_id, is_public
        FROM pastes
        WHERE id = $1
    `
	row := config.DB.QueryRow(query, pasteID)
	err := row.Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.Language,
		&res.Password,
		&res.BurnAfterRead,
		&res.ExpiresAt,
		&res.CreatedAt,
		&res.Views,
		&res.UserID,
		&res.IsPublic,
	)
	if err != nil {
		return res, err
	}

	_, err = config.DB.Exec("UPDATE pastes SET views = views + 1 WHERE id = $1", pasteID)
	if err != nil {
		return res, fmt.Errorf("failed to increment views: %v", err)
	}

	if res.BurnAfterRead {
		_, _ = config.DB.Exec("DELETE FROM pastes WHERE id = $1", pasteID)
	}

	return res, nil
}
