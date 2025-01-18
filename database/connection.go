package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetConnection(dbUrl string) (context.Context, *pgx.Conn, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbUrl)

	return ctx, conn, err
}
