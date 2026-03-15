-- name: GetAllCock :many
SELECT * 
FROM cocktails
where is_available = true
order by is_new desc, name;
