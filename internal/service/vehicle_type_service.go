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

type VehicleTypeService interface {
	AllVehicleType(ctx context.Context, req request.BasePagination) response.BaseResponsePagination[response.VehicleType]
	VehicleTypeById(ctx context.Context, code string) response.VehicleType
	CreateVehicleType(ctx context.Context, req request.CreateVehicleTypeRequest)
	UpdateVehicleType(ctx context.Context, req request.UpdateVehicleTypeRequest)
}

type vehicleTypeService struct {
	repo      repository.VehicleTypeRepository
	exception exception.Exception
	conf      config.Config
}

func NewVehicleTypeService(repo repository.VehicleTypeRepository) VehicleTypeService {
	return &vehicleTypeService{
		repo:      repo,
		exception: exception.NewException("vehicleType-service"),
		conf:      *config.Get(),
	}
}

func (s *vehicleTypeService) AllVehicleType(ctx context.Context, req request.BasePagination) response.BaseResponsePagination[response.VehicleType] {
	data, err := s.repo.AllVehicleType(ctx, req)
	s.exception.PanicIfError(err, true)
	s.exception.IsNotFound(data, true)

	content := model.ContentPagination[response.VehicleType]{
		Count:   data.Count,
		Content: []response.VehicleType{},
	}

	for _, val := range data.Content {
		res := response.VehicleType{
			Code: val.Code,
			Name: val.Name,
		}
		res.SetPrice(val.Price)
		content.Content = append(content.Content, res)
	}

	return response.WithPagination[response.VehicleType](content, req)
}

func (s *vehicleTypeService) VehicleTypeById(ctx context.Context, code string) response.VehicleType {
	data, err := s.repo.VehicleTypeByCode(ctx, code)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFound(data, false)

	res := response.VehicleType{
		Code: data.Code,
		Name: data.Name,
	}
	res.SetPrice(data.Price)
	return res
}

func (s *vehicleTypeService) handleErrorUniqueVehicleType(err error, isList bool) {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return
	}

	if pgErr.Code != pgerrcode.UniqueViolation {
		return
	}

	switch pgErr.ConstraintName {
	case "vehicle_type_pkey":
		s.exception.IsBadRequestMessage("Code already exist.", isList)
	}
}

func (s *vehicleTypeService) CreateVehicleType(ctx context.Context, req request.CreateVehicleTypeRequest) {
	payload := sql.CreateVehicleTypeParams{
		Code:      req.Code,
		Name:      req.Name,
		Price:     utils.PGNumericFloat64(req.Price),
		CreatedBy: req.UserId,
	}

	err := s.repo.CreateVehicleType(ctx, payload)
	s.handleErrorUniqueVehicleType(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *vehicleTypeService) UpdateVehicleType(ctx context.Context, req request.UpdateVehicleTypeRequest) {
	payload := sql.UpdateVehicleTypeParams{
		Code:      req.Code,
		Name:      req.Name,
		Price:     utils.PGNumericFloat64(req.Price),
		UpdatedBy: utils.PGInt32(req.UserId),
	}

	err := s.repo.UpdateVehicleType(ctx, payload)
	s.handleErrorUniqueVehicleType(err, false)
	s.exception.PanicIfError(err, false)
}
