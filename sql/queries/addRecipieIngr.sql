-- name: AddRecipIngr :one
INSERT INTO recipies(cocktail_id, ingredient_id, amount, unit, created_at, modified_at)
VALUES (
	$1,
	$2,
	$3,
	$4,
	now(),
	now()
)
RETURNING *;
