-- +goose Up
CREATE TABLE users(
	id UUID UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP NOT NULL,
	last_seen_at TIMESTAMP NOT NULL
);

-- +goose Down
drop table users;
