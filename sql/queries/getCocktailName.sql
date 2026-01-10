-- name: GetCocktailName :one
SELECT * FROM cocktails
WHERE id = $1;
