// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: vehicle_type.sql

package sql

import (
	"context"
)

const countAllVehicleType = `-- name: CountAllVehicleType :one
SELECT COUNT(*) FROM vehicle_type
`

func (q *Queries) CountAllVehicleType(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAllVehicleType)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAllVehicleType = `-- name: GetAllVehicleType :many
SELECT code, name, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM vehicle_type
`

func (q *Queries) GetAllVehicleType(ctx context.Context) ([]VehicleType, error) {
	rows, err := q.db.Query(ctx, getAllVehicleType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VehicleType{}
	for rows.Next() {
		var i VehicleType
		if err := rows.Scan(
			&i.Code,
			&i.Name,
			&i.Price,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.DeletedAt,
			&i.DeletedBy,
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

const getVehicleTypeByCode = `-- name: GetVehicleTypeByCode :one
SELECT code, name, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by FROM vehicle_type where code = $1 LIMIT 1
`

func (q *Queries) GetVehicleTypeByCode(ctx context.Context, code string) (VehicleType, error) {
	row := q.db.QueryRow(ctx, getVehicleTypeByCode, code)
	var i VehicleType
	err := row.Scan(
		&i.Code,
		&i.Name,
		&i.Price,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.DeletedAt,
		&i.DeletedBy,
	)
	return i, err
}
