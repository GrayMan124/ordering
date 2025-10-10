-- name: SendOrder :one
INSERT INTO orders (id, created_at, modified_at, cocktail_id, ordered_by, canceled_at, finished)
VALUES (
	gen_random_uuid(),
	now(),
	now(),
	$1,
	$2,
	NULL,
	false
)
RETURNING *;
