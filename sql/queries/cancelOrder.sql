-- name: CancelOrder :one
UPDATE orders
SET canceled_at = now(), modified_at = now()
where id = $1
RETURNING *;
