-- name: GetMyOrders :many
select ord.id, ord.created_at, ord.canceled_at, ord.finished,  cock.name, cock.img_name
from orders ord
left join cocktails cock on cock.id = ord.cocktail_id
where ord.ordered_by = $1
order by ord.created_at desc;
