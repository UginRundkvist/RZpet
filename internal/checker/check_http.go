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
			Status:       false,
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
		Status:       true,
		ResponseCode: resp.StatusCode,
		ErrorText:    "",
		DurationMs:   elapsed.Milliseconds(),
	}
}

func CheckHtttp(t task.TaskRequest) {
	go func(t task.TaskRequest) {
		for t.Active {
			//res := CheckURL(t.Target)

			time.Sleep(t.Interval)
		}
	}(t)
}
