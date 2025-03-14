CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(100) UNIQUE   NOT NULL,
    password   VARCHAR(100)          NOT NULL,
    is_active  BOOLEAN DEFAULT false NOT NULL,
    created_at TIMESTAMP             NOT NULL,
    updated_at TIMESTAMP             NULL
);