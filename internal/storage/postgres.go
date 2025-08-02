package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func PgConnect(ctx context.Context, connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// defer conn.Close(ctx) Вынести в функцию меин

	return conn, nil
}