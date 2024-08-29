-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';




CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY ASC,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at INTEGER,
    updated_at INTEGER
);




-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';




DROP TABLE IF EXISTS users;




-- +goose StatementEnd
