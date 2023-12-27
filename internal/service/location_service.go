package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/model"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"be-park-ease/internal/sql"
	"be-park-ease/utils"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type LocationService interface {
	AllLocation(ctx context.Context, req request.BasePagination) response.BaseResponsePagination[response.Location]
	LocationByCode(ctx context.Context, code string) response.Location
	CreateLocation(ctx context.Context, req request.CreateLocationRequest)
	UpdateLocation(ctx context.Context, req request.UpdateLocationRequest)
	DeleteLocation(ctx context.Context, req request.DeleteLocationRequest)
}

type locationService struct {
	repo      repository.LocationRepository
	exception exception.Exception
	conf      config.Config
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{
		repo:      repo,
		exception: exception.NewException("location-service"),
		conf:      *config.Get(),
	}
}

func (l *locationService) AllLocation(ctx context.Context, req request.BasePagination) response.BaseResponsePagination[response.Location] {
	data, err := l.repo.AllLocation(ctx, req)
	l.exception.PanicIfError(err, true)
	l.exception.IsNotFound(data, true)

	content := model.ContentPagination[response.Location]{
		Count:   data.Count,
		Content: []response.Location{},
	}

	for _, val := range data.Content {
		res := response.Location{
			Code:   val.Code,
			Name:   val.Name,
			IsExit: val.IsExit,
		}
		content.Content = append(content.Content, res)
	}

	return response.WithPagination[response.Location](content, req)
}

func (l *locationService) LocationByCode(ctx context.Context, code string) response.Location {
	data, err := l.repo.LocationByCode(ctx, code)
	l.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	l.exception.IsNotFound(data, false)

	res := response.Location{
		Code:   data.Code,
		Name:   data.Name,
		IsExit: data.IsExit,
	}

	return res
}

func (l *locationService) CreateLocation(ctx context.Context, req request.CreateLocationRequest) {
	payload := sql.CreateLocationParams{
		Name:      req.Name,
		IsExit:    req.IsExit,
		CreatedBy: req.UserId,
	}

	err := l.repo.CreateLocation(ctx, payload)
	l.exception.PanicIfError(err, false)
	l.handleErrorUniqueLocation(err, false)
}

func (l *locationService) UpdateLocation(ctx context.Context, req request.UpdateLocationRequest) {
	payload := sql.UpdateLocationParams{
		Code:      req.Code,
		Name:      req.Name,
		IsExit:    req.IsExit,
		UpdatedBy: utils.PGInt32(req.UserId),
	}

	err := l.repo.UpdateLocation(ctx, payload)
	l.handleErrorUniqueLocation(err, false)
	l.exception.PanicIfError(err, false)
}

func (l *locationService) DeleteLocation(ctx context.Context, req request.DeleteLocationRequest) {
	payload := sql.DeleteLocationParams{
		Code:      req.Code,
		DeletedBy: utils.PGInt32(req.UserId),
	}

	err := l.repo.DeleteLocation(ctx, payload)
	l.exception.PanicIfError(err, false)
	l.handleErrorUniqueLocation(err, false)
}

func (l *locationService) handleErrorUniqueLocation(err error, isList bool) {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return
	}

	if pgErr.Code != pgerrcode.UniqueViolation {
		return
	}

	switch pgErr.ConstraintName {
	case "location_pkey":
		l.exception.IsBadRequestMessage("Location already exist.", isList)
	}
}
