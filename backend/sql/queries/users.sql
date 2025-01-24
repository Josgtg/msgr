-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

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

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING id;
