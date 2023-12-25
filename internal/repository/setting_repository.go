package repository

import (
	"be-park-ease/internal/model"
	"be-park-ease/internal/sql"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SettingRepository interface {
	GetAllSetting(ctx context.Context) ([]sql.Setting, error)
	CreateOrUpdateSetting(ctx context.Context, req []model.CreateOrUpdateSetting) error
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

func (r *settingRepository) CreateOrUpdateSetting(ctx context.Context, req []model.CreateOrUpdateSetting) error {
	var args []any
	query := `INSERT INTO setting (key, value)
        VALUES %s ON CONFLICT ON CONSTRAINT setting_pkey DO UPDATE SET
        value = EXCLUDED.value,
        description = EXCLUDED.description`

	value := ""
	for idx, setting := range req {
		currentIdx := idx*2 + 1
		value += fmt.Sprintf("($%d, $%d)", currentIdx, currentIdx+1)
		args = append(args, []any{setting.Key, setting.Value}...)
		if idx != len(req)-1 {
			value += ","
		}
	}

	_, err := r.db.Exec(ctx, fmt.Sprintf(query, value), args...)
	return err
}
