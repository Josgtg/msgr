-- name: GetAllChats :many
SELECT * FROM chats;

-- name: GetChat :one
SELECT * FROM chats
WHERE id = $1;

-- name: GetChatsByUsers :many
SELECT * FROM chats
WHERE first_user = $1 OR second_user = $1;

-- name: ChatExists :one
SELECT EXISTS (
    SELECT TRUE FROM chats
    WHERE id = $1 LIMIT 1
);

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

-- name: DeleteChat :one
DELETE FROM chats
WHERE id = $1
RETURNING id;
