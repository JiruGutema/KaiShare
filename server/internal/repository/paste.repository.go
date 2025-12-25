// Package repository interacts with database
package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jirugutema/kaishare/internal/config"
	"github.com/jirugutema/kaishare/internal/dto"
)

func CreatePaste(paste dto.PasteDTO) error {
	query := `
        INSERT INTO pastes (
            id, title, content, language, password, burn_after_read, expires_at, created_at, views, user_id, is_public
        ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
    `
	_, err := config.DB.Exec(
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
	return err
}

func GetPaste(pasteID uuid.UUID) (dto.PasteDTO, error) {
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
	return res, err
}

func IncrementViews(pasteID uuid.UUID) error {
	_, err := config.DB.Exec("UPDATE pastes SET views = views + 1 WHERE id = $1", pasteID)
	return err
}

func DeletePaste(pasteID uuid.UUID) error {
	_, err := config.DB.Exec("DELETE FROM pastes WHERE id = $1", pasteID)
	return err
}

func PasteUserExists(userID uuid.UUID) (bool, error) {
	var exists bool
	err := config.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)", userID).Scan(&exists)
	return exists, err
}

func GetMyPastes(userID uuid.UUID) (dto.MyPastesDTO, error) {
	var myPastes dto.MyPastesDTO
	query := "SELECT id, title, content, user_id, created_at, expires_at, language, burn_after_read, views, is_public  FROM pastes where user_id = $1"
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return myPastes, err
	}

	for rows.Next() {
		var paste dto.PasteResponse
		err := rows.Scan(
			&paste.ID,
			&paste.Title,
			&paste.Content,
			&paste.UserID,
			&paste.CreatedAt,
			&paste.ExpiresAt,
			&paste.Language,
			&paste.BurnAfterRead,
			&paste.Views,
			&paste.IsPublic,
		)
		if err != nil {
			return myPastes, err
		}

		myPastes = append(myPastes, paste)
	}
	defer rows.Close()
	return myPastes, err
}

func UpdatePaste(paste dto.UpdatePasteDTO) (*dto.PasteResponse, error) {
	if paste.ID == uuid.Nil {
		return nil, errors.New("missing paste ID")
	}

	fields := []string{}
	args := []any{}
	i := 1

	fmt.Print(fields)

	if paste.Title != nil {
		fields = append(fields, fmt.Sprintf("title = $%d", i))
		args = append(args, *paste.Title)
		i++
	}

	if paste.Content != nil {
		fields = append(fields, fmt.Sprintf("content = $%d", i))
		args = append(args, *paste.Content)
		i++
	}

	if paste.Language != nil {
		fields = append(fields, fmt.Sprintf("language = $%d", i))
		args = append(args, *paste.Language)
		i++
	}

	if paste.BurnAfterRead != nil {
		fields = append(fields, fmt.Sprintf("burn_after_read = $%d", i))
		args = append(args, *paste.BurnAfterRead)
		i++
	}

	if paste.ExpiresAt != nil {
		fields = append(fields, fmt.Sprintf("expires_at = $%d", i))
		args = append(args, *paste.ExpiresAt)
		i++
	}

	if paste.IsPublic != nil {
		fields = append(fields, fmt.Sprintf("is_public = $%d", i))
		args = append(args, *paste.IsPublic)
		i++
	}

	if len(fields) == 0 {
		return nil, errors.New("no fields to update")
	}

	// WHERE id
	args = append(args, paste.ID)

	query := fmt.Sprintf(`
		UPDATE pastes
		SET %s
		WHERE id = $%d
		RETURNING
			id,
			title,
			content,
			language,
			burn_after_read,
			expires_at,
			is_public,
			created_at,
	`,
		strings.Join(fields, ", "),
		i,
	)

	var res dto.PasteResponse

	err := config.DB.QueryRow(query, args...).Scan(
		&res.ID,
		&res.Title,
		&res.Content,
		&res.Language,
		&res.BurnAfterRead,
		&res.ExpiresAt,
		&res.IsPublic,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
