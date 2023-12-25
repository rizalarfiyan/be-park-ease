package service

import (
	"be-park-ease/config"
	"be-park-ease/exception"
	"be-park-ease/internal/repository"
	"be-park-ease/internal/response"
	"context"
)

type SettingService interface {
	GetAllSetting(ctx context.Context) response.SettingResponse
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
