-- +goose Up
CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    lyrics TEXT NOT NULL,
    link VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS songs;