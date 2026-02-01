-- name: GetAllCock :many
SELECT * FROM cocktails
order by is_new desc, name;
