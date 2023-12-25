package repository

import (
	"be-park-ease/internal/sql"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SettingRepository interface {
	GetAllSetting(ctx context.Context) ([]sql.Setting, error)
}

type settingRepository struct {
	db    *pgxpool.Pool
	query *sql.Queries
}

func NewSettingRepository(db *pgxpool.Pool) SettingRepository {
	return &settingRepository{
		db:    db,
		query: sql.New(db),
	}
}

func (r *settingRepository) GetAllSetting(ctx context.Context) ([]sql.Setting, error) {
	return r.query.GetAllSetting(ctx)
}
