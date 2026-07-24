-- name: GetAllCock :many
SELECT cock.* 
FROM cocktails cock
where cock.is_available = true
order by cock.is_new cock.desc, cock.name;
