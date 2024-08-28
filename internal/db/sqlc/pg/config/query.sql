-- name: FindUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1
LIMIT
    1;




-- name: InsertUser :one
INSERT INTO
    users (name, email, password)
VALUES
    ($1, $2, $3) RETURNING *;
