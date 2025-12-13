package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"titile"`
	Content      string    `json:"content"`
	UserID       uuid.UUID `json:"userId"`
	Read         bool      `json:"read"`
	RelationLink string    `json:"relationLink"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
