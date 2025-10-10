-- +goose Up
ALTER TABLE cocktails 
ADD COLUMN base_spirit TEXT NOT NULL DEFAULT 'unset';
ALTER TABLE cocktails 
ADD COLUMN cocktail_type TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
ALTER TABLE cocktails 
DROP COLUMN base_spirit;
ALTER TABLE cocktails 
DROP COLUMN cocktail_type;
