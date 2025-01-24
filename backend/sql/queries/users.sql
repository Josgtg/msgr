-- name: GetAllUsers :many
SELECT id, name, email, registered_at FROM users;

-- name: GetUser :one
SELECT id, name, email, registered_at FROM users
WHERE id = $1;

-- name: UserExists :one
SELECT EXISTS (
    SELECT TRUE FROM users
    WHERE id = $1 LIMIT 1
);

-- name: IsUsedEmail :one
SELECT EXISTS (
    SELECT TRUE FROM users
    WHERE email = $1 LIMIT 1
);

-- name: Login :one
SELECT id, name, email, registered_at FROM users
WHERE email = $1 AND password = $2 LIMIT 1;

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
RETURNING (id, name, email, registered_at);

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
