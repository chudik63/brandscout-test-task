package models

import "time"

type Quote struct {
	ID        uint64    `json:"id"`
	Author    string    `json:"author"`
	Quote     string    `json:"quote"`
	CreatedAt time.Time `json:"created_at"`
}
