-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';




CREATE TABLE users (
    id BIGINT PRIMARY KEY NOT NULL GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NULL DEFAULT now(),
    updated_at TIMESTAMP NULL DEFAULT now()
);




CREATE UNIQUE INDEX master_unique_email ON users (email);




-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';




DROP TABLE IF EXISTS users;




-- +goose StatementEnd
