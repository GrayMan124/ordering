-- name: GetCocktail :one
SELECT * FROM cocktails
WHERE id = $1;
