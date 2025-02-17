-- +goose Up
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short TEXT UNIQUE NOT NULL,
    original TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE urls;