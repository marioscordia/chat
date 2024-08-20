package model

import "time"

// Chat is an object with chat information
type Chat struct {
	ID           int64
	Name         string
	CreatorID    int64
	MessageCount int64
	Type         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// ChatCreate is an object with creating parameters
type ChatCreate struct {
	Name      string
	CreatorID int64
	UserIDs   []int64
	Type      string
}
