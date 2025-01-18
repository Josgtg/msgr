-- name: GetAllUsers :many
SELECT * FROM users;

-- name: InsertUser :one
INSERT INTO users(
    id,
    name,
    password,
    email
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
