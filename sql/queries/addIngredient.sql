-- name: AddIngredient :one
INSERT INTO ingredients(id, name, abv, is_available, created_at, modified_at) 
VALUES (
	gen_random_uuid(),
	$1,
	$2,
	TRUE,
	now(),
	now()
)
RETURNING *;
