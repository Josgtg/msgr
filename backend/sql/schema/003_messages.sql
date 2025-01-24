-- +goose Up

CREATE TABLE messages (
    id UUID PRIMARY KEY NOT NULL,
    chat UUID NOT NULL REFERENCES chats(id),
    sender UUID NOT NULL REFERENCES users(id),
    message TEXT NOT NULL,
    sent_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

DROP TABLE messages;
