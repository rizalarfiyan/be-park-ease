// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: history.sql

package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countAllHistory = `-- name: CountAllHistory :one
select COUNT(*) from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
`

func (q *Queries) CountAllHistory(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAllHistory)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createEntryHistory = `-- name: CreateEntryHistory :exec
INSERT INTO entry_history (id, location_code, vehicle_type_code, vehicle_number, created_by)
VALUES ($1, $2, $3, $4, $5)
`

type CreateEntryHistoryParams struct {
	ID              interface{}
	LocationCode    string
	VehicleTypeCode string
	VehicleNumber   string
	CreatedBy       int32
}

func (q *Queries) CreateEntryHistory(ctx context.Context, arg CreateEntryHistoryParams) error {
	_, err := q.db.Exec(ctx, createEntryHistory,
		arg.ID,
		arg.LocationCode,
		arg.VehicleTypeCode,
		arg.VehicleNumber,
		arg.CreatedBy,
	)
	return err
}

const createExitHistory = `-- name: CreateExitHistory :exec
INSERT INTO exit_history (entry_history_id, location_code, price, exited_by)
VALUES ($1, $2, $3, $4)
`

type CreateExitHistoryParams struct {
	EntryHistoryID string
	LocationCode   string
	Price          pgtype.Numeric
	ExitedBy       int32
}

func (q *Queries) CreateExitHistory(ctx context.Context, arg CreateExitHistoryParams) error {
	_, err := q.db.Exec(ctx, createExitHistory,
		arg.EntryHistoryID,
		arg.LocationCode,
		arg.Price,
		arg.ExitedBy,
	)
	return err
}

const createFineHistory = `-- name: CreateFineHistory :exec
INSERT INTO fine_history (entry_history_id, location_code, price, identity, vehicle_identity, name, address, description, fined_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateFineHistoryParams struct {
	EntryHistoryID  string
	LocationCode    string
	Price           pgtype.Numeric
	Identity        string
	VehicleIdentity string
	Name            string
	Address         string
	Description     string
	FinedBy         int32
}

func (q *Queries) CreateFineHistory(ctx context.Context, arg CreateFineHistoryParams) error {
	_, err := q.db.Exec(ctx, createFineHistory,
		arg.EntryHistoryID,
		arg.LocationCode,
		arg.Price,
		arg.Identity,
		arg.VehicleIdentity,
		arg.Name,
		arg.Address,
		arg.Description,
		arg.FinedBy,
	)
	return err
}

const getAllHistory = `-- name: GetAllHistory :many
select CAST(eh.id as varchar) as id, coalesce(fh.location_code, coalesce(exh.location_code, eh.location_code)) as location_code, eh.vehicle_type_code, eh.vehicle_number, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) as date,
       CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
`

type GetAllHistoryRow struct {
	ID              string
	LocationCode    string
	VehicleTypeCode string
	VehicleNumber   string
	Date            pgtype.Timestamp
	Type            string
}

func (q *Queries) GetAllHistory(ctx context.Context) ([]GetAllHistoryRow, error) {
	rows, err := q.db.Query(ctx, getAllHistory)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllHistoryRow{}
	for rows.Next() {
		var i GetAllHistoryRow
		if err := rows.Scan(
			&i.ID,
			&i.LocationCode,
			&i.VehicleTypeCode,
			&i.VehicleNumber,
			&i.Date,
			&i.Type,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllHistoryStatistic = `-- name: GetAllHistoryStatistic :many
select MIN(coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)))::date as date, SUM(coalesce(fh.price, coalesce(exh.price, 0))) as revenue, COUNT(*) as vehicle
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
`

type GetAllHistoryStatisticRow struct {
	Date    pgtype.Date
	Revenue int64
	Vehicle int64
}

func (q *Queries) GetAllHistoryStatistic(ctx context.Context) ([]GetAllHistoryStatisticRow, error) {
	rows, err := q.db.Query(ctx, getAllHistoryStatistic)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllHistoryStatisticRow{}
	for rows.Next() {
		var i GetAllHistoryStatisticRow
		if err := rows.Scan(&i.Date, &i.Revenue, &i.Vehicle); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountHistoryStatistic = `-- name: GetCountHistoryStatistic :one
select 
    count(*) as total,
    CAST(COALESCE(SUM(coalesce(fh.price, coalesce(exh.price, 0))), 0) AS float) as revenue,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NULL AND fh.fined_at IS NULL THEN 1 ELSE 0 END), 0) AS float) AS entry_total,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NOT NULL THEN 1 ELSE 0 END), 0) AS int) AS exit_total,
    CAST(COALESCE(SUM(CASE WHEN exh.exited_at IS NOT NULL THEN exh.price ELSE 0 END), 0) AS float) AS exit_revenue,
    CAST(COALESCE(SUM(CASE WHEN fh.fined_at IS NOT NULL THEN 1 ELSE 0 END), 0) AS int) AS fine_total,
    CAST(COALESCE(SUM(CASE WHEN fh.fined_at IS NOT NULL THEN fh.price ELSE 0 END), 0) AS float) AS fine_revenue
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
WHERE coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) BETWEEN $1::timestamp AND $2::timestamp
LIMIT 1
`

type GetCountHistoryStatisticParams struct {
	StartAt pgtype.Timestamp
	EndAt   pgtype.Timestamp
}

type GetCountHistoryStatisticRow struct {
	Total       int64
	Revenue     float64
	EntryTotal  float64
	ExitTotal   int32
	ExitRevenue float64
	FineTotal   int32
	FineRevenue float64
}

func (q *Queries) GetCountHistoryStatistic(ctx context.Context, arg GetCountHistoryStatisticParams) (GetCountHistoryStatisticRow, error) {
	row := q.db.QueryRow(ctx, getCountHistoryStatistic, arg.StartAt, arg.EndAt)
	var i GetCountHistoryStatisticRow
	err := row.Scan(
		&i.Total,
		&i.Revenue,
		&i.EntryTotal,
		&i.ExitTotal,
		&i.ExitRevenue,
		&i.FineTotal,
		&i.FineRevenue,
	)
	return i, err
}

const getDataByEntryHistoryId = `-- name: GetDataByEntryHistoryId :one
select eh.id, vt.price, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) date,
   CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
JOIN vehicle_type vt on eh.vehicle_type_code = vt.code
where eh.id = $1
ORDER BY date DESC
LIMIT 1
`

type GetDataByEntryHistoryIdRow struct {
	ID    interface{}
	Price pgtype.Numeric
	Date  pgtype.Timestamp
	Type  string
}

func (q *Queries) GetDataByEntryHistoryId(ctx context.Context, id interface{}) (GetDataByEntryHistoryIdRow, error) {
	row := q.db.QueryRow(ctx, getDataByEntryHistoryId, id)
	var i GetDataByEntryHistoryIdRow
	err := row.Scan(
		&i.ID,
		&i.Price,
		&i.Date,
		&i.Type,
	)
	return i, err
}

const getLastHistoryByVehicleNumber = `-- name: GetLastHistoryByVehicleNumber :one
select eh.vehicle_number, coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) date,
       CASE WHEN fh.fined_at IS NOT NULL THEN 'fine' WHEN exh.exited_at IS NOT NULL THEN 'exit' ELSE 'entry' END AS type
from entry_history eh
LEFT JOIN exit_history exh on eh.id = exh.entry_history_id
LEFT JOIN fine_history fh on eh.id = fh.entry_history_id
WHERE eh.vehicle_number = $1
ORDER BY date DESC
LIMIT 1
`

type GetLastHistoryByVehicleNumberRow struct {
	VehicleNumber string
	Date          pgtype.Timestamp
	Type          string
}

func (q *Queries) GetLastHistoryByVehicleNumber(ctx context.Context, vehicleNumber string) (GetLastHistoryByVehicleNumberRow, error) {
	row := q.db.QueryRow(ctx, getLastHistoryByVehicleNumber, vehicleNumber)
	var i GetLastHistoryByVehicleNumberRow
	err := row.Scan(&i.VehicleNumber, &i.Date, &i.Type)
	return i, err
}
