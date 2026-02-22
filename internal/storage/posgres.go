package storage

import (
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
		tassk_id SERIAL PRIMARY KEY,
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

// func SaveTaskResult(pool *pgxpool.Pool, result checker.TaskResult) error {
// 	query := `
//     INSERT INTO task_results (task_name, status, response)
//     VALUES ($1, $2, $3)
//     `
// 	_, err := pool.Exec(context.Background(),
// 		query,
// 		result.TaskName,
// 		result.Status,
// 		result.Response,
// 	)
// 	return err
// }
