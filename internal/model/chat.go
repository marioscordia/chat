package model

import "time"

// Chat is an object with chat information
type Chat struct {
	ID            int64
	Name          string
	CreatorID     int64
	MesseageCount int64
	Type          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
