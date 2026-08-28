-- name: GetCockSearch :many
SELECT cock.*, x.score
FROM cocktails cock
  left join (
  select cast(sum(ingr.abv * rec.amount) / 100 as DOUBLE PRECISION)as score, cock.name 
  from cocktails cock
  left join recipies rec on rec.cocktail_id = cock.id
  left join ingredients ingr on ingr.id = rec.ingredient_id 
  group by cock.name
  ) x on x.name = cock.name
WHERE cock.name like '%' || $1 || '%'
order by cock.is_new desc, cock.name;
