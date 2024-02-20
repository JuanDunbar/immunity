package events

import "time"

type SuspiciousActivity struct {
	Event     string    `json:"event"`
	Data      string    `json:"data"`
	Rule      int       `json:"rule"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
}
