package models

import "time"

type Rule struct {
	ID          string
	Query       string
	Description string
	Action      string
	LastUsed    time.Time
}
