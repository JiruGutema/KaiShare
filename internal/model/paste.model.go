package model

import "time"

type Paste struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	Language      string     `json:"language"`
	Password      *string    `json:"password"`
	BurnAfterRead bool       `json:"burnAfterRead"`
	ExpiresAt     *time.Time `json:"expiresAt"`
	CreatedAt     time.Time  `json:"createdAt"`
	Views         time.Time  `json:"views"`
	UserID        *string    `json:"userId"`
	IsPublic      bool       `json:"isPublic"`
}
