-- +goose Up
CREATE TABLE cocktails(
	id UUID UNIQUE,
	created_at TIMESTAMP,
	updated_at TIMESTAMP,
	data_url TEXT
);
-- +goose Down
DROP TABLE cocktails;

