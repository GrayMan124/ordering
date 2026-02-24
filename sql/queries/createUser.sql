-- name: CreateUser :one
INSERT INTO users(id, name, created_at, modified_at, last_seen_at)
VALUES (
	gen_random_uuid(),
	$1,
	now(),
	now(),
	now()
)
RETURNING *;
