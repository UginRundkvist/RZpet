package checker

type TaskResult struct {
	Date         string // дата записи
	Status       bool   // "success - true" или "error - false"
	ResponseCode int    // статус
	DurationMs   int64  // продолжительность
	ErrorText    string // текст ошибки
}
