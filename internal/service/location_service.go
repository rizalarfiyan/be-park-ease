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
	AllLocation(ctx context.Context, req request.GetAllLocationRequest) response.BaseResponsePagination[response.Location]
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

func (s *locationService) AllLocation(ctx context.Context, req request.GetAllLocationRequest) response.BaseResponsePagination[response.Location] {
	data, err := s.repo.AllLocation(ctx, req)
	s.exception.PanicIfError(err, true)
	s.exception.IsNotFound(data, true)

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

	return response.WithPagination[response.Location](content, req.BasePagination)
}

func (s *locationService) LocationByCode(ctx context.Context, code string) response.Location {
	data, err := s.repo.LocationByCode(ctx, code)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFound(data, false)

	res := response.Location{
		Code:   data.Code,
		Name:   data.Name,
		IsExit: data.IsExit,
	}

	return res
}

func (s *locationService) CreateLocation(ctx context.Context, req request.CreateLocationRequest) {
	payload := sql.CreateLocationParams{
		Code:      req.Code,
		Name:      req.Name,
		IsExit:    req.IsExit,
		CreatedBy: req.UserId,
	}

	err := s.repo.CreateLocation(ctx, payload)
	s.handleErrorUniqueLocation(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *locationService) UpdateLocation(ctx context.Context, req request.UpdateLocationRequest) {
	payload := sql.UpdateLocationParams{
		Code:      req.Code,
		Name:      req.Name,
		IsExit:    req.IsExit,
		UpdatedBy: utils.PGInt32(req.UserId),
	}

	err := s.repo.UpdateLocation(ctx, payload)
	s.handleErrorUniqueLocation(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *locationService) DeleteLocation(ctx context.Context, req request.DeleteLocationRequest) {
	payload := sql.DeleteLocationParams{
		Code:      req.Code,
		DeletedBy: utils.PGInt32(req.UserId),
	}

	err := s.repo.DeleteLocation(ctx, payload)
	s.handleErrorUniqueLocation(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *locationService) handleErrorUniqueLocation(err error, isList bool) {
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
		s.exception.IsBadRequestMessage("Location already exist.", isList)
	}
}
