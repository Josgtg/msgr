// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: chats.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const deleteChat = `-- name: DeleteChat :one
DELETE FROM chats
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteChat(ctx context.Context, id pgtype.UUID) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, deleteChat, id)
	err := row.Scan(&id)
	return id, err
}

const getAllChats = `-- name: GetAllChats :many
SELECT id, first_user, second_user, created_at FROM chats
`

func (q *Queries) GetAllChats(ctx context.Context) ([]Chat, error) {
	rows, err := q.db.Query(ctx, getAllChats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Chat
	for rows.Next() {
		var i Chat
		if err := rows.Scan(
			&i.ID,
			&i.FirstUser,
			&i.SecondUser,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChat = `-- name: GetChat :one
SELECT id, first_user, second_user, created_at FROM chats
WHERE id = $1
`

func (q *Queries) GetChat(ctx context.Context, id pgtype.UUID) (Chat, error) {
	row := q.db.QueryRow(ctx, getChat, id)
	var i Chat
	err := row.Scan(
		&i.ID,
		&i.FirstUser,
		&i.SecondUser,
		&i.CreatedAt,
	)
	return i, err
}

const insertChat = `-- name: InsertChat :one
INSERT INTO chats(
    id,
	first_user,
	second_user
) VALUES (
    $1,
    $2,
    $3
)
RETURNING id
`

type InsertChatParams struct {
	ID         pgtype.UUID
	FirstUser  pgtype.UUID
	SecondUser pgtype.UUID
}

func (q *Queries) InsertChat(ctx context.Context, arg InsertChatParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, insertChat, arg.ID, arg.FirstUser, arg.SecondUser)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}
