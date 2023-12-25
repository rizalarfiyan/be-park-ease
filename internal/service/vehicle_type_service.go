package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/model"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"context"
)

type VehicleTypeService interface {
	AllVehicleType(ctx context.Context, req request.BasePagination) response.BaseResponsePagination[response.VehicleType]
	VehicleTypeById(ctx context.Context, code string) response.VehicleType
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
