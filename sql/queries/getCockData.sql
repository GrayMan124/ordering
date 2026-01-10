-- name: GetCocktail :one
SELECT * FROM cocktails
WHERE name = $1;
