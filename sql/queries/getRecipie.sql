-- name: GetRecipie :many
SELECT ing.* 
FROM ingredients ing  
INNER JOIN cocktails cock on cock.ID = ing.cocktail_ID
WHERE $1 = cock.Name;
