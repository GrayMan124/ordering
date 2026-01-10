-- +goose Up
CREATE TABLE ingredients(
	ID UUID UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	quantity INTEGER NOT NULL, 
	abv REAL NOT NULL,
	created_at TIMESTAMP NOT NULL, 
	modified_at TIMESTAMP,
	cocktail_id UUID NOT NULL REFERENCES cocktails(id)
);

-- +goose Down
DROP TABLE ingredients;
