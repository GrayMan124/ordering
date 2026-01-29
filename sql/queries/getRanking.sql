-- name: GetRanking :many
select sum(ingr.abv * ingr.quantity) / 100 as score, ord.ordered_by
from orders ord
left join ingredients ingr on ingr.cocktail_id = ord.cocktail_id
where ord.finished = true
group by ord.ordered_by
order by score desc;
