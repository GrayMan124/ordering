-- name: GetFullOrders :many
SELECT  ord.id, 
	ord.ordered_by,
	ord.created_at,
	ord.cocktail_id,
	ing.id,
	ing.name,
	ing.quantity,
	ing.abv     
FROM orders ord 
LEFT OUTER JOIN ingredients ing on ing.cocktail_id = ord.cocktail_id
WHERE ord.canceled_at is null
and ord.finished = false;
