package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type pgConnector struct {
	Conn *pgx.Conn
}

func NewPostgres(ctx context.Context, connString string) (*pgConnector, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pg: %w", err)
	}
	return &pgConnector{Conn: conn}, nil
}

func (p *pgConnector) Close(ctx context.Context) error {
	return p.Conn.Close(ctx)
}