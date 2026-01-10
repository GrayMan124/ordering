-- +goose Up
ALTER TABLE cocktails
ADD COLUMN img_name TEXT;

-- +goose Down
ALTER TABLE cocktails
DROP COLUMN img_name;
