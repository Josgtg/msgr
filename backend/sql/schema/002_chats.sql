-- +goose Up

CREATE TABLE chats (
    id UUID PRIMARY KEY NOT NULL,
    first_user UUID NOT NULL REFERENCES users(id),
    second_user UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE chats;
