package models

import "time"

type Quote struct {
	ID        uint64
	Author    string
	Quote     string
	CreatedAt time.Time
}
