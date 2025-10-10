-- name: AddCocktail :one
INSERT INTO cocktails (id, created_at, updated_at, data_url, base_spirit, cocktail_type, name) 
VALUES (
	gen_random_uuid(),
	now(),
	now(),
	$1,
	$2,
	$3,
	$4
)
RETURNING *;
