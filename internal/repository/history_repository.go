package repository

import (
	"be-park-ease/constants"
	"be-park-ease/internal/model"
	"be-park-ease/internal/request"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HistoryRepository interface {
	GetAllHistory(ctx context.Context, req request.AllHistoryRequest) (*model.ContentPagination[sql.GetAllHistoryRow], error)
	GetLastHistoryByVehicleNumber(ctx context.Context, vehicleNumber string) (sql.GetLastHistoryByVehicleNumberRow, error)
	GetDataByEntryHistoryId(ctx context.Context, entryHistoryId string) (sql.GetDataByEntryHistoryIdRow, error)
	CreateEntryHistory(ctx context.Context, req sql.CreateEntryHistoryParams) error
	CreateExitHistory(ctx context.Context, req sql.CreateExitHistoryParams) error
	CreateFineHistory(ctx context.Context, req sql.CreateFineHistoryParams) error
	GetCountHistoryStatistic(ctx context.Context, req sql.GetCountHistoryStatisticParams) (sql.GetCountHistoryStatisticRow, error)
	AllHistoryStatistic(ctx context.Context, req model.AllHistoryStatistic) ([]sql.GetAllHistoryStatisticRow, error)
}

type historyRepository struct {
	db           *pgxpool.Pool
	query        *sql.Queries
	queryBuilder *sql.Queries
}

func NewHistoryRepository(db *pgxpool.Pool) HistoryRepository {
	return &historyRepository{
		db:           db,
		query:        sql.New(db),
		queryBuilder: sql.New(utils.QueryWrap(db)),
	}
}

func (r historyRepository) GetAllHistory(ctx context.Context, req request.AllHistoryRequest) (*model.ContentPagination[sql.GetAllHistoryRow], error) {
	var res model.ContentPagination[sql.GetAllHistoryRow]

	baseBuilder := func(b *utils.QueryBuilder) {
		switch req.HistoryType {
		case constants.HistoryTypeEntry:
			b.Where("exh.exited_at IS NULL AND fh.fined_at IS NULL")
		case constants.HistoryTypeExit:
			b.Where("exh.exited_at IS NOT NULL")
		case constants.HistoryTypeFine:
			b.Where("fh.fined_at IS NOT NULL")
		}

		if req.VehicleType != "" {
			b.Where("eh.vehicle_type_code = $1", req.VehicleType)
		}

		if req.Location != "" {
			b.Where("eh.location_code = $1", req.Location)
		}

		if req.Search != "" {
			b.Where("LOWER(eh.id) LIKE $1 OR LOWER(eh.vehicle_number) LIKE $1", fmt.Sprintf("%%%s%%", req.Search))
		}
	}

	entryHistory, err := r.queryBuilder.GetAllHistory(utils.QueryBuild(ctx, func(b *utils.QueryBuilder) {
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

	count, err := r.queryBuilder.CountAllHistory(utils.QueryBuild(ctx, baseBuilder))
	if err != nil {
		return nil, err
	}

	res.Content = entryHistory
	res.Count = count
	return &res, nil
}

func (r historyRepository) GetLastHistoryByVehicleNumber(ctx context.Context, vehicleNumber string) (sql.GetLastHistoryByVehicleNumberRow, error) {
	return r.query.GetLastHistoryByVehicleNumber(ctx, vehicleNumber)
}

func (r historyRepository) GetDataByEntryHistoryId(ctx context.Context, entryHistoryId string) (sql.GetDataByEntryHistoryIdRow, error) {
	return r.query.GetDataByEntryHistoryId(ctx, entryHistoryId)
}

func (r historyRepository) CreateEntryHistory(ctx context.Context, req sql.CreateEntryHistoryParams) error {
	return r.query.CreateEntryHistory(ctx, req)
}

func (r historyRepository) CreateExitHistory(ctx context.Context, req sql.CreateExitHistoryParams) error {
	return r.query.CreateExitHistory(ctx, req)
}

func (r historyRepository) CreateFineHistory(ctx context.Context, req sql.CreateFineHistoryParams) error {
	return r.query.CreateFineHistory(ctx, req)
}

func (r historyRepository) GetCountHistoryStatistic(ctx context.Context, req sql.GetCountHistoryStatisticParams) (sql.GetCountHistoryStatisticRow, error) {
	return r.query.GetCountHistoryStatistic(ctx, req)
}

func (r *historyRepository) AllHistoryStatistic(ctx context.Context, req model.AllHistoryStatistic) ([]sql.GetAllHistoryStatisticRow, error) {
	histories, err := r.queryBuilder.GetAllHistoryStatistic(utils.QueryBuild(ctx, func(b *utils.QueryBuilder) {
		b.Where("coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)) BETWEEN $1 AND $2", req.StartDate, req.EndDate)

		switch req.TimeFrequency {
		case constants.FilterTimeFrequencyToday:
			b.GroupBy("DATE_TRUNC('hour', coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)))")
		case constants.FilterTimeFrequencyYear:
			b.GroupBy("DATE_TRUNC('month', coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)))")
		default:
			b.GroupBy("DATE_TRUNC('day', coalesce(fh.fined_at, coalesce(exh.exited_at, eh.created_at)))")
		}

		b.Order("date DESC")
	}))

	if err != nil {
		return nil, err
	}

	return histories, nil
}
