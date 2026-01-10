-- name: AddRecipie :one
INSERT INTO ingredients(id, name, quantity, abv, created_at, modified_at, cocktail_id) 
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	$3,
	now(),
	now(),
	$4
)
RETURNING *;
