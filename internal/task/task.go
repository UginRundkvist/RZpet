package task

import (
	"time"
)

type TaskRequest struct {
	ID       int
	UserID   int
	Type     string // "http", "ping", "tcp"
	Target   string // нет в бд
	Interval time.Duration
	Active   bool
	Payload  string // вид задачи
}
