-- name: FinishOrder :one
UPDATE orders
SET finished = TRUE, modified_at = now()
where id = $1
RETURNING *;
