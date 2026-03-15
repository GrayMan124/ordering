-- name: GetCocksIngr :many
select cock.name 
from ingredients ingr
left join recipies rec on rec.ingredient_id = ingr.id
left join cocktails cock on cock.id = rec.cocktail_id
where ingr.name = $1; 

