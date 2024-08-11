package model

import "time"

// Message is an object with message information
type Message struct {
	ID        int64
	ChatID    int64
	UserID    int64
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
