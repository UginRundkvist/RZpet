package checker

type TaskResult struct {
	Status       string // "success" или "error"
	ResponseCode int
	DurationMs   int64
	ErrorText    string
}
