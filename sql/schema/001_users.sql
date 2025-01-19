-- +goose Up

CREATE TABLE users (
    id UUID UNIQUE PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    registered_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE users;
