package main

import (
	"automation-service/internal/checker"
	"automation-service/internal/storage"
	"fmt"
	"log"
	"time"
)

func main() {
	dsn := "postgres://admin:admin@localhost:5432/automation?sslmode=disable"

	pool, err := storage.New(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = storage.CreateSchema(pool)
	if err != nil {
		fmt.Println("err Create")
	}

	// Создаём дату
	tam := time.Date(2026, time.February, 23, 14, 30, 0, 0, time.UTC)

	// Преобразуем в строку для TIMESTAMP
	dateStr := tam.Format("2006-01-02 15:04:05")

	test_task := checker.TaskResult{
		Date:         dateStr,
		Status:       true,
		ResponseCode: 123,
		ErrorText:    "",
		DurationMs:   203,
	}

	err = storage.SaveTaskResult(pool, test_task)
	if err != nil {
		fmt.Println(err)
	}

	defer pool.Close()

	fmt.Println("Connected to database successfully")

	// пока просто держим программу живой
	select {}
}
