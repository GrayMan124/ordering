-- name: GetFullOrders :many
SELECT  ord.id, 
	ord.ordered_by,
	ord.created_at,
	ord.cocktail_id,
	cock.name
FROM orders ord 
LEFT OUTER JOIN cocktails cock on cock.id = ord.cocktail_id
WHERE ord.canceled_at is null
and ord.finished = false;
