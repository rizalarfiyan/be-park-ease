// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: user.sql

package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countAllUser = `-- name: CountAllUser :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountAllUser(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countAllUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAllUser = `-- name: GetAllUser :many
select id, name, username, password, role, status, token, expired_at, created_at, updated_at from users
`

func (q *Queries) GetAllUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Username,
			&i.Password,
			&i.Role,
			&i.Status,
			&i.Token,
			&i.ExpiredAt,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const getUserByToken = `-- name: GetUserByToken :one
SELECT id, name, username, password, role, status, token, expired_at, created_at, updated_at FROM users WHERE token = $1 LIMIT 1
`

func (q *Queries) GetUserByToken(ctx context.Context, token pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUserByToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, name, username, password, role, status, token, expired_at, created_at, updated_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Username,
		&i.Password,
		&i.Role,
		&i.Status,
		&i.Token,
		&i.ExpiredAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserToken = `-- name: UpdateUserToken :exec
UPDATE users
SET token = $1, expired_at = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3
`

type UpdateUserTokenParams struct {
	Token     pgtype.Text
	ExpiredAt pgtype.Timestamp
	ID        int32
}

func (q *Queries) UpdateUserToken(ctx context.Context, arg UpdateUserTokenParams) error {
	_, err := q.db.Exec(ctx, updateUserToken, arg.Token, arg.ExpiredAt, arg.ID)
	return err
}
