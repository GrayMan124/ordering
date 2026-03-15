-- name: GetCockSearch :many
SELECT * 
FROM cocktails
WHERE name like '%' || $1 || '%'
order by is_new desc, name;
