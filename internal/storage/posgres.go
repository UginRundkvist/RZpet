package storage

import (
	"automation-service/internal/checker"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("cannot create pool: %w", err)
	}

	// проверяем соединение
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("cannot ping db: %w", err)
	}

	return pool, nil
}

func CreateSchema(pool *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS task_results (
		task_id SERIAL PRIMARY KEY,
		date TIMESTAMP NOT NULL,
		status BOOLEAN NOT NULL,
		response_code INTEGER NOT NULL,
		duration int NOT NULL,
		error_text VARCHAR(255)
	);
	`
	_, err := pool.Exec(context.Background(), query)
	return err
}

func SaveTaskResult(pool *pgxpool.Pool, t checker.TaskResult) error {
	query := `
        INSERT INTO task_results (task_id, date, status, response_code, duration, error_text)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	ctx := context.Background()
	//id := uuid.New()
	_, err := pool.Exec(
		ctx,
		query,
		13123131,
		t.Date,
		t.Status,
		t.ResponseCode,
		t.DurationMs,
		t.ErrorText,
	)

	if err != nil {
		return fmt.Errorf("failed to save task result: %w", err)
	}

	return nil
}
