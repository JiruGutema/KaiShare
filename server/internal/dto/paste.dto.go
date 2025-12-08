package dto

import (
	"time"

	"github.com/google/uuid"
)

type PasteDTO struct {
	ID            uuid.UUID  `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	Language      string     `json:"language"`
	Password      *string    `json:"password"`
	BurnAfterRead bool       `json:"burnAfterRead"`
	ExpiresAt     *time.Time `json:"expiresAt"`
	CreatedAt     time.Time  `json:"createdAt"`
	Views         int        `json:"views"`
	UserID        *uuid.UUID `json:"userId"`
	IsPublic      bool       `json:"isPublic"`
}

type PasteResponse struct {
	ID               uuid.UUID  `json:"id"`
	Title            string     `json:"title"`
	Content          string     `json:"content"`
	Language         string     `json:"language"`
	BurnAfterRead    bool       `json:"burnAfterRead"`
	ExpiresAt        *time.Time `json:"expiresAt"`
	CreatedAt        time.Time  `json:"createdAt"`
	Views            int        `json:"views"`
	UserID           *uuid.UUID `json:"userId"`
	IsPublic         bool       `json:"isPublic"`
	RequiresPassword bool       `json:"requiresPassword"`
}

type MyPastesDTO []PasteDTO
