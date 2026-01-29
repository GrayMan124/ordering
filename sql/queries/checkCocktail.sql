-- name: CheckCocktail :one
SELECT count(*) FROM cocktails
WHERE name = $1;
