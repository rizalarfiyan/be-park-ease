package service

import (
	"be-park-ease/config"
	"be-park-ease/constants"
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
	"strings"
	"time"
)

type HistoryService interface {
	AllHistory(ctx context.Context, req request.AllHistoryRequest) response.BaseResponsePagination[response.EntryHistory]
	CreateEntryHistory(ctx context.Context, req request.CreateEntryHistoryRequest)
}

type historyService struct {
	repo         repository.HistoryRepository
	repoLocation repository.LocationRepository
	exception    exception.Exception
	conf         config.Config
}

func NewHistoryService(repo repository.HistoryRepository, repoLocation repository.LocationRepository) HistoryService {
	return &historyService{
		repo:         repo,
		repoLocation: repoLocation,
		exception:    exception.NewException("history-service"),
		conf:         *config.Get(),
	}
}

func (s *historyService) AllHistory(ctx context.Context, req request.AllHistoryRequest) response.BaseResponsePagination[response.EntryHistory] {
	data, err := s.repo.GetAllHistory(ctx, req)
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

func (s *historyService) handleErrorEntryHistory(err error, isList bool) {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return
	}

	if pgErr.Code != pgerrcode.ForeignKeyViolation {
		return
	}

	switch pgErr.ConstraintName {
	case "entry_history_location_code_fkey":
		s.exception.IsBadRequestMessage("Location is not available.", isList)
	case "entry_history_vehicle_type_code_fkey":
		s.exception.IsBadRequestMessage("Vehicle Type is not available.", isList)
	}
}

func (s *historyService) CreateEntryHistory(ctx context.Context, req request.CreateEntryHistoryRequest) {
	lastHistory, err := s.repo.GetLastHistoryByVehicleNumber(ctx, req.VehicleNumber)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	if !utils.IsEmpty(lastHistory) && strings.EqualFold(lastHistory.Type, string(constants.HistoryTypeEntry)) {
		s.exception.IsBadRequestMessage("Vehicle already entry, please check your ticket to exit", false)
	}

	location, err := s.repoLocation.LocationByCode(ctx, req.LocationCode)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFoundMessage(location, "Location is not available.", false)
	if location.IsExit {
		s.exception.IsBadRequestMessage("Please use a entry location", false)
	}

	payload := sql.CreateEntryHistoryParams{
		ID:              utils.GenerateEntryHistoryId(),
		LocationCode:    req.LocationCode,
		VehicleTypeCode: req.VehicleTypeCode,
		VehicleNumber:   req.VehicleNumber,
		CreatedBy:       req.UserId,
	}

	err = s.repo.CreateEntryHistory(ctx, payload)
	s.handleErrorEntryHistory(err, false)
	s.exception.PanicIfError(err, false)
}
