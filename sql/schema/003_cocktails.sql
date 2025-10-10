-- +goose Up
ALTER TABLE cocktails
ADD COLUMN name TEXT NOT NULL UNIQUE;

-- +goose Down
ALTER TABLE cocktails
DROP COLUMN name;
