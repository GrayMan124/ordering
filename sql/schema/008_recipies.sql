-- sqlc: ignore	
-- +goose Up
DROP TABLE IF EXISTS ingredients;
CREATE TABLE ingredients(
	ID UUID UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	abv REAL NOT NULL,
	is_available BOOLEAN NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP
);
CREATE TABLE recipies(
	cocktail_id UUID NOT NULL REFERENCES cocktails(id) ON DELETE CASCADE,
	ingredient_id UUID NOT NULL REFERENCES ingredients(id) ON DELETE CASCADE,
	amount INTEGER NOT NULL,
	unit TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
	modified_at TIMESTAMP,
	PRIMARY KEY (cocktail_id, ingredient_id)
);

-- sqlc: ignore	
-- +goose Down
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS recipies;
CREATE TABLE ingredients(
	ID UUID UNIQUE PRIMARY KEY,
	name TEXT NOT NULL,
	quantity INTEGER NOT NULL, 
	abv REAL NOT NULL,
	created_at TIMESTAMP NOT NULL, 
	modified_at TIMESTAMP,
	cocktail_id UUID NOT NULL REFERENCES cocktails(id)
);

