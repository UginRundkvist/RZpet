package task

import (
	"time"
)

type Task struct {
	ID       int
	UserID   int
	Type     string // "http", "ping", "tcp"
	Target   string // URL, IP, host
	Interval time.Duration
	Active   bool
	Payload  string // JSON со специфичными параметрами
}
