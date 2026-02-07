-- name: GetIngredients :many
select id, name
from ingredients
where name = ANY($1::text[]);
