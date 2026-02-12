-- name: GetRanking :many
select cast(sum(ingr.abv * rec.amount) / 100 as DOUBLE PRECISION)as score, ord.ordered_by
from orders ord
left join recipies rec on rec.cocktail_id = ord.cocktail_id
left join ingredients ingr on ingr.id = rec.ingredient_id
where ord.finished = true
group by ord.ordered_by
order by score desc;
