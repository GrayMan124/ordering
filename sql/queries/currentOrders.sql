-- name: GetCurrentOrders :many
SELECT * 
FROM orders 
WHERE canceled_at is null
and finished = false;
