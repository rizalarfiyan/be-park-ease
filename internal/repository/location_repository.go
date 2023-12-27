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

type LocationRepository interface {
	AllLocation(ctx context.Context, req request.BasePagination) (*model.ContentPagination[sql.Location], error)
	LocationByCode(ctx context.Context, code string) (sql.Location, error)
	CreateLocation(ctx context.Context, req sql.CreateLocationParams) error
	UpdateLocation(ctx context.Context, req sql.UpdateLocationParams) error
	DeleteLocation(ctx context.Context, req sql.DeleteLocationParams) error
}

type locationRepository struct {
	db           *pgxpool.Pool
	query        *sql.Queries
	queryBuilder *sql.Queries
}

func NewLocationRepository(db *pgxpool.Pool) LocationRepository {
	return &locationRepository{
		db:           db,
		query:        sql.New(db),
		queryBuilder: sql.New(db),
	}
}

func (l locationRepository) AllLocation(ctx context.Context, req request.BasePagination) (*model.ContentPagination[sql.Location], error) {
	var res model.ContentPagination[sql.Location]

	baseBuilder := func(b *utils.QueryBuilder) {
		b.Where("deleted_at IS NULL")

		if req.Search != "" {
			b.Where("LOWER(name) LIKE $1 OR LOWER(code) LIKE $1", fmt.Sprintf("%%%s%%", req.Search))
		}
	}

	entryHistory, err := l.queryBuilder.GetAllLocation(utils.QueryBuild(ctx, func(b *utils.QueryBuilder) {
		baseBuilder(b)
		if req.OrderBy != "" && req.Order != "" {
			b.Ordering(req.OrderBy, req.Order)
		} else {
			b.Order("date DESC")
		}
		b.Pagination(req.Page, req.Limit)
	}))
	if err != nil {
		return nil, err
	}

	count, err := l.queryBuilder.CountAllLocation(utils.QueryBuild(ctx, baseBuilder))
	if err != nil {
		return nil, err
	}

	res.Content = entryHistory
	res.Count = count
	return &res, nil
}

func (l locationRepository) LocationByCode(ctx context.Context, code string) (sql.Location, error) {
	return l.query.GetLocationByCode(ctx, code)
}

func (l locationRepository) CreateLocation(ctx context.Context, req sql.CreateLocationParams) error {
	return l.query.CreateLocation(ctx, req)
}

func (l locationRepository) UpdateLocation(ctx context.Context, req sql.UpdateLocationParams) error {
	return l.query.UpdateLocation(ctx, req)
}

func (l locationRepository) DeleteLocation(ctx context.Context, req sql.DeleteLocationParams) error {
	return l.query.DeleteLocation(ctx, req)
}