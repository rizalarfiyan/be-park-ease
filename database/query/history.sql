-- name: GetAllHistory :many
select eh.id, eh.location_code, eh.vehicle_type_code, eh.vehicle_number, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) date,
       CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id;

-- name: CountAllHistory :one
select COUNT(*) from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id;
