-- name: AddCocktail :one
INSERT INTO cocktails (id, created_at, updated_at, data_url, base_spirit, cocktail_type, name, img_name, type, is_new) 
VALUES (
	gen_random_uuid(),
	now(),
	now(),
	Null,
	$1,
	$2,
	$3,
	$4,
	$5,
	false
)
RETURNING *;
