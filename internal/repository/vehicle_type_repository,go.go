package repository

import (
	"be-park-ease/internal/model"
	"be-park-ease/internal/request"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleTypeRepository interface {
	AllVehicleType(ctx context.Context, req request.BasePagination) (*model.ContentPagination[sql.VehicleType], error)
}

type vehicleTypeRepository struct {
	db           *pgxpool.Pool
	query        *sql.Queries
	queryBuilder *sql.Queries
}

func NewVehicleTypeRepository(db *pgxpool.Pool) VehicleTypeRepository {
	return &vehicleTypeRepository{
		db:           db,
		query:        sql.New(db),
		queryBuilder: sql.New(utils.QueryWrap(db)),
	}
}

func (r *vehicleTypeRepository) AllVehicleType(ctx context.Context, req request.BasePagination) (*model.ContentPagination[sql.VehicleType], error) {
	var res model.ContentPagination[sql.VehicleType]

	baseBuilder := func(b *utils.QueryBuilder) {
		b.Where("deleted_at IS NULL")

		if req.Search != "" {
			b.Where("LOWER(name) LIKE $1 OR LOWER(code) LIKE $1", fmt.Sprintf("%%%s%%", req.Search))
		}
	}

	entryHistory, err := r.queryBuilder.GetAllVehicleType(utils.QueryBuild(ctx, func(b *utils.QueryBuilder) {
		baseBuilder(b)
		if req.OrderBy != "" && req.Order != "" {
			b.Ordering(req.OrderBy, req.Order)
		} else {
			b.Order("created_at DESC")
		}
		b.Pagination(req.Page, req.Limit)
	}))
	if err != nil {
		return nil, err
	}

	count, err := r.queryBuilder.CountAllVehicleType(utils.QueryBuild(ctx, baseBuilder))
	if err != nil {
		return nil, err
	}

	res.Content = entryHistory
	res.Count = count
	return &res, nil
}
