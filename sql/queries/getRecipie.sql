-- name: GetRecipie :many
select ingr.name
from cocktails cock
left join recipies rec on rec.cocktail_id = cock.id
left join ingredients ingr on ingr.id = rec.ingredient_id
where cock.name = $1 
and ingr.id is not null
.
