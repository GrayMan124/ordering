-- name: UpdateAvCock :exec
UPDATE cocktails 
SET is_available = true, updated_at = now()
where cocktails.id not in (
  select cock.id
  from cocktails cock 
  left join recipies rec on rec.cocktail_id = cock.id 
  left join ingredients ingr on ingr.id = rec.ingredient_id
  where ingr.is_available = false);
