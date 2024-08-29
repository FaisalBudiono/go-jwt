CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY ASC,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at INTEGER,
    updated_at INTEGER
);
