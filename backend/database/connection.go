package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnection(dbUrl string) (context.Context, *pgxpool.Pool, error) {
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dbUrl)

	return ctx, conn, err
}
