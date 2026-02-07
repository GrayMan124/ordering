-- +goose Up
ALTER TABLE cocktails
ADD COLUMN type TEXT,
ADD COLUMN is_new boolean NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE cocktails
DROP COLUMN type,
DROP COLUMN is_new;
