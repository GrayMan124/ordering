-- name: GetFullRecipie :many
select ingr.name, rc.amount, rc.unit
from recipies rc
left join ingredients ingr on rc.ingredient_id = ingr.id
where rc.cocktail_id = $1;
