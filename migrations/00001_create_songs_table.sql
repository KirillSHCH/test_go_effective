-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    group VARCHAR(255) NOT NULL DEFAULT '',
    song VARCHAR(255) NOT NULL DEFAULT '',
    text TEXT NOT NULL DEFAULT '',
    link VARCHAR(255) NOT NULL DEFAULT '',
    release_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books
-- +goose StatementEnd
