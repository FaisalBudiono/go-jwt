-- name: FindUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = ?
LIMIT
    1;




-- name: AllUsers :many
SELECT
    *
FROM
    users;




-- name: InsertUser :one
INSERT INTO
    users (name, email, password, created_at, updated_at)
VALUES
    (?, ?, ?, ?, ?) RETURNING *;




-- name: TruncateUsers :exec
DELETE FROM users;
