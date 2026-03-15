-- name: UpdateIngr :one
UPDATE ingredients 
SET is_available =$1 , modified_at = now()
where name = $2
RETURNING *;
