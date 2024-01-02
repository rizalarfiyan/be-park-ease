-- name: GetAllHistory :many
select CAST(eh.id as varchar) as id, coalesce(fh.location_code, coalesce(exh.location_code, eh.location_code)) as location_code, eh.vehicle_type_code, eh.vehicle_number, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) as date,
       CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id;

-- name: CountAllHistory :one
select COUNT(*) from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id;

-- name: GetLastHistoryByVehicleNumber :one
select eh.vehicle_number, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) date,
       CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
WHERE eh.vehicle_number = $1
ORDER BY date DESC
LIMIT 1;

-- name: GetDataByEntryHistoryId :one
select eh.id, vt.price, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) date,
   CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
JOIN vehicle_type vt on eh.vehicle_type_code = vt.code
where eh.id = $1
ORDER BY date DESC
LIMIT 1;

-- name: CreateEntryHistory :exec
INSERT INTO entry_history (id, location_code, vehicle_type_code, vehicle_number, created_by)
VALUES ($1, $2, $3, $4, $5);

-- name: CreateExitHistory :exec
INSERT INTO exit_history (entry_history_id, location_code, price, exited_by)
VALUES ($1, $2, $3, $4);

-- name: CreateFineHistory :exec
INSERT INTO fine_history (entry_history_id, location_code, price, identity, vehicle_identity, name, address, description, fined_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetAllHistoryStatistic :many
select MIN(coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)))::date as date, SUM(coalesce(fh.price, coalesce(exh.price, 0))) as revenue, COUNT(*) as vehicle
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id;

-- name: GetCountHistoryStatistic :one
select 
    count(*) as total,
    CAST(COALESCE(SUM(coalesce(fh.price, coalesce(exh.price, 0))), 0) AS float) as revenue,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NULL AND fh.fined_at IS NULL THEN 1 ELSE 0 END), 0) AS float) AS entry_total,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NOT NULL THEN 1 ELSE 0 END), 0) AS int) AS exit_total,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NOT NULL THEN exh.price ELSE 0 END), 0) AS float) AS exit_revenue,
    CAST(COALESCE(SUM(CASE WHEN fh.fined_at IS NOT NULL THEN 1 ELSE 0 END), 0) AS int) AS fine_total,
    CAST(COALESCE(SUM(CASE WHEN fh.fined_at IS NOT NULL THEN exh.price ELSE 0 END), 0) AS float) AS fine_revenue
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
WHERE coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) BETWEEN sqlc.arg(start_at)::timestamp AND sqlc.arg(end_at)::timestamp
LIMIT 1;
