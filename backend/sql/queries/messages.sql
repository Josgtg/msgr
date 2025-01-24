-- name: GetAllMessages :many
SELECT * FROM messages;

-- name: GetMessage :one
SELECT * FROM messages
WHERE id = $1;

-- name: GetMessagesByChat :many
SELECT * FROM messages
WHERE chat = $1;

-- name: MessageExists :one
SELECT EXISTS (
    SELECT TRUE FROM messages
    WHERE id = $1 LIMIT 1
);

-- name: InsertMessage :one
INSERT INTO messages(
    id,
    chat,
    sender,
    message
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteMessage :one
DELETE FROM messages
WHERE id = $1
RETURNING id;
