package checker

type TaskResult struct {
	Status       bool   // "success - true" или "error - false"
	ResponseCode int    // статус
	DurationMs   int64  // продолжительность
	ErrorText    string // текст ошибки
	Date         string // дата записи
}
