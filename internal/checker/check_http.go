package checker

import (
	"automation-service/internal/task"
	"fmt"
	"net/http"
	"time"
)

func CheckURL(url string) TaskResult {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return TaskResult{
			Status:       "error",
			ResponseCode: 0,
			ErrorText:    err.Error(),
			DurationMs:   time.Since(start).Milliseconds(),
		}
	}
	defer resp.Body.Close()
	elapsed := time.Since(start)

	fmt.Println(elapsed)
	fmt.Println("Status", resp.StatusCode)
	return TaskResult{
		Status:       "success",
		ResponseCode: resp.StatusCode,
		ErrorText:    "",
		DurationMs:   elapsed.Milliseconds(),
	}
}

func CheckHtttp(t task.Task) {
	go func(t task.Task) {

		time.Sleep(t.Interval)
	}(t)
}
