-- +goose Up
ALTER TABLE cocktails
ADD COLUMN is_mocktail BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE cocktails
ADD COLUMN is_available BOOLEAN NOT NULL DEFAULT TRUE;

-- +goose Down
ALTER TABLE cocktails
DROP COLUMN is_mocktail ;

ALTER TABLE cocktails
DROP COLUMN is_available;
