package repository

import (
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	GetAllByToken(ctx context.Context, token string) (sql.User, error)
	GetAllByUsername(ctx context.Context, username string) (sql.User, error)
	UpdateUserToken(ctx context.Context, arg sql.UpdateUserTokenParams) error
}

type authRepository struct {
	db    *pgxpool.Pool
	query *sql.Queries
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{
		db:    db,
		query: sql.New(db),
	}
}

func (r authRepository) GetAllByToken(ctx context.Context, token string) (sql.User, error) {
	return r.query.GetAllByToken(ctx, utils.PGText(token))
}

func (r authRepository) GetAllByUsername(ctx context.Context, username string) (sql.User, error) {
	return r.query.GetAllByUsername(ctx, username)
}

func (r authRepository) UpdateUserToken(ctx context.Context, arg sql.UpdateUserTokenParams) error {
	return r.query.UpdateUserToken(ctx, arg)
}
