package service

import (
	"be-park-ease/config"
	"be-park-ease/constants"
	"be-park-ease/exception"
	"be-park-ease/internal/model"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/request"
	"be-park-ease/internal/response"
	"context"
	"fmt"
)

type SettingService interface {
	GetAllSetting(ctx context.Context) response.SettingResponse
	CreateOrUpdateSetting(ctx context.Context, req request.CreateOrUpdateSettingRequest)
}

type settingService struct {
	repo      repository.SettingRepository
	exception exception.Exception
	conf      config.Config
}

func NewSettingService(repo repository.SettingRepository) SettingService {
	return &settingService{
		repo:      repo,
		exception: exception.NewException("setting-service"),
		conf:      *config.Get(),
	}
}

func (s *settingService) GetAllSetting(ctx context.Context) response.SettingResponse {
	settings, err := s.repo.GetAllSetting(ctx)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)

	res := response.SettingResponse{}
	res.FromDB(settings)
	return res
}

func (s *settingService) CreateOrUpdateSetting(ctx context.Context, req request.CreateOrUpdateSettingRequest) {
	payload := []model.CreateOrUpdateSetting{
		{
			Key:   constants.SettingFineTicketCalculation,
			Value: fmt.Sprint(req.FineTicketCalculation),
		},
		{
			Key:   constants.SettingNextHourCalculation,
			Value: fmt.Sprint(req.NextHourCalculation),
		},
		{
			Key:   constants.SettingMaxCapacity,
			Value: fmt.Sprint(req.MaxCapacity),
		},
	}
	err := s.repo.CreateOrUpdateSetting(ctx, payload)
	s.exception.PanicIfError(err, false)
}
