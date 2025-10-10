-- +goose Up
CREATE TABLE orders (
	ID UUID UNIQUE PRIMARY KEY,
	created_at TIMESTAMP NOT NULL, 
	modified_at TIMESTAMP,
	cocktail_id UUID NOT NULL REFERENCES cocktails(id),
	ordered_by TEXT NOT NULL,
	canceled_at TIMESTAMP,
	finished BOOLEAN DEFAULT false
);

-- +goose Down
DROP TABLE orders;
