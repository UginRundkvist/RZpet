package task

import (
	"time"
)

type Task struct {
	ID       int
	UserID   int
	Type     string // "http", "ping", "tcp"
	Target   string // ссылка на сайт
	Interval time.Duration
	Active   bool
	Payload  string // вид задачи
}
