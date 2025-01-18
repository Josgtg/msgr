-- name: GetAllChats :many
SELECT * FROM chats;

-- name: InsertChat :one
INSERT INTO chats(
    id,
	first_user,
	second_user
) VALUES (
    $1,
    $2,
    $3
)
RETURNING id;

-- name: GetChat :one
SELECT * FROM chats
WHERE id = $1;

-- name: DeleteChat :one
DELETE FROM chats
WHERE id = $1
RETURNING id;
