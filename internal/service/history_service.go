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
	"time"
)

type HistoryService interface {
	AllHistory(ctx context.Context, req request.AllHistoryRequest) response.BaseResponsePagination[response.EntryHistory]
}

type historyService struct {
	repo      repository.HistoryRepository
	exception exception.Exception
	conf      config.Config
}

func NewHistoryService(repo repository.HistoryRepository) HistoryService {
	return &historyService{
		repo:      repo,
		exception: exception.NewException("history-service"),
		conf:      *config.Get(),
	}
}

func (s *historyService) AllHistory(ctx context.Context, req request.AllHistoryRequest) response.BaseResponsePagination[response.EntryHistory] {
	data, err := s.repo.GetAllEntryHistory(ctx, req)
	s.exception.PanicIfError(err, true)
	s.exception.IsNotFound(data, true)

	content := model.ContentPagination[response.EntryHistory]{
		Count:   data.Count,
		Content: []response.EntryHistory{},
	}

	for _, val := range data.Content {
		res := response.EntryHistory{
			ID:              val.ID,
			LocationCode:    val.LocationCode,
			VehicleTypeCode: val.VehicleTypeCode,
			VehicleNumber:   val.VehicleNumber,
			Type:            constants.HistoryType(val.Type),
		}
		if val.Date.Valid {
			res.Date = val.Date.Time.Format(time.DateTime)
		}
		if !res.Type.IsValid() {
			res.Type = constants.HistoryTypeEntry
		}
		content.Content = append(content.Content, res)
	}

	return response.WithPagination[response.EntryHistory](content, req.BasePagination)
}
