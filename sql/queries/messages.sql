-- name: GetAllMessages :many
SELECT * FROM messages;

-- name: InsertMessage :one
INSERT INTO messages(
    id,
    chat,
    sender,
    receiver,
    message
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING id;

-- name: GetMessage :one
SELECT * FROM messages
WHERE id = $1;

-- name: DeleteMessage :one
DELETE FROM messages
WHERE id = $1
RETURNING id;
