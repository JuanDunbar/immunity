package events

import "time"

type SuspiciousActivity struct {
	Event     string    `json:"event"`
	Data      string    `json:"data"`
	Rule      string    `json:"rule"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
}
