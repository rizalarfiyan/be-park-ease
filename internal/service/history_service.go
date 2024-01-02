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
	"math"
	"strings"
	"time"
)

type HistoryService interface {
	AllHistory(ctx context.Context, req request.AllHistoryRequest) response.BaseResponsePagination[response.EntryHistory]
	CreateEntryHistory(ctx context.Context, req request.CreateEntryHistoryRequest)
	CalculatePriceHistory(ctx context.Context, req request.CalculatePriceHistoryRequest) float64
	CreateExitHistory(ctx context.Context, req request.CreateExitHistoryRequest)
	CreateFineHistory(ctx context.Context, req request.CreateFineHistoryRequest)
	GetAllHistoryStatistic(ctx context.Context, req request.TimeFrequency) response.HistoryStatistic
}

type historyService struct {
	repo           repository.HistoryRepository
	repoLocation   repository.LocationRepository
	exception      exception.Exception
	conf           config.Config
	serviceSetting SettingService
}

func NewHistoryService(repo repository.HistoryRepository, repoLocation repository.LocationRepository, serviceSetting SettingService) HistoryService {
	return &historyService{
		repo:           repo,
		repoLocation:   repoLocation,
		exception:      exception.NewException("history-service"),
		conf:           *config.Get(),
		serviceSetting: serviceSetting,
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
		s.exception.IsUnprocessableEntityMessage("Vehicle already entry, please check your ticket to exit", false)
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

func (s *historyService) CalculatePriceHistory(ctx context.Context, req request.CalculatePriceHistoryRequest) float64 {
	entryHistory, err := s.repo.GetDataByEntryHistoryId(ctx, req.EntryHistoryId)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFoundMessage(entryHistory, "Entry History is not available.", false)

	if !utils.IsEmpty(entryHistory) && !strings.EqualFold(entryHistory.Type, string(constants.HistoryTypeEntry)) {
		s.exception.IsUnprocessableEntityMessage("Vehicle already exit or fine, please entry first for vehicle", false)
	}

	setting := s.serviceSetting.GetAllSetting(ctx)
	var vehicleTypePrice float64
	rawPrice, err := entryHistory.Price.Float64Value()
	if entryHistory.Price.Valid && err == nil {
		vehicleTypePrice = rawPrice.Float64
	}

	now := time.Now()
	entryTime := now
	if entryHistory.Date.Valid {
		entryTime = entryHistory.Date.Time
	}

	price := 1 * vehicleTypePrice
	nextHours := math.Ceil(now.Sub(entryTime).Hours()) - 1
	if nextHours > 0 {
		price += nextHours * (float64(setting.NextHourCalculation) / 100) * vehicleTypePrice
	}

	if req.IsFine {
		price += float64(setting.FineTicketCalculation) / 100 * vehicleTypePrice
	}

	return price
}

func (s *historyService) handleErrorExitHistory(err error, isList bool) {
	var pgErr *pgconn.PgError
	ok := errors.As(err, &pgErr)
	if !ok {
		return
	}

	if pgErr.Code != pgerrcode.ForeignKeyViolation {
		return
	}

	switch pgErr.ConstraintName {
	case "exit_history_location_code_fkey":
		s.exception.IsBadRequestMessage("Location is not available.", isList)
	}
}

func (s *historyService) ValidateExitFineHistory(ctx context.Context, locationCode, entryHistoryId string) {
	location, err := s.repoLocation.LocationByCode(ctx, locationCode)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFoundMessage(location, "Location is not available.", false)
	if !location.IsExit {
		s.exception.IsBadRequestMessage("Please use a exit location", false)
	}

	entryHistory, err := s.repo.GetDataByEntryHistoryId(ctx, entryHistoryId)
	s.exception.PanicIfErrorWithoutNoSqlResult(err, false)
	s.exception.IsNotFoundMessage(entryHistory, "Entry History is not available.", false)

	if !strings.EqualFold(entryHistory.Type, string(constants.HistoryTypeEntry)) {
		s.exception.IsUnprocessableEntityMessage("Vehicle already exit or fine, please entry first for vehicle", false)
	}
}

func (s *historyService) CreateExitHistory(ctx context.Context, req request.CreateExitHistoryRequest) {
	s.ValidateExitFineHistory(ctx, req.LocationCode, req.EntryHistoryId)

	payload := sql.CreateExitHistoryParams{
		EntryHistoryID: req.EntryHistoryId,
		LocationCode:   req.LocationCode,
		Price:          utils.PGNumericFloat64(req.Price),
		ExitedBy:       req.UserId,
	}

	err := s.repo.CreateExitHistory(ctx, payload)
	s.handleErrorExitHistory(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *historyService) CreateFineHistory(ctx context.Context, req request.CreateFineHistoryRequest) {
	s.ValidateExitFineHistory(ctx, req.LocationCode, req.EntryHistoryId)

	payload := sql.CreateFineHistoryParams{
		EntryHistoryID:  req.EntryHistoryId,
		LocationCode:    req.LocationCode,
		Price:           utils.PGNumericFloat64(req.Price),
		Identity:        req.Identity,
		VehicleIdentity: req.VehicleIdentity,
		Name:            req.Name,
		Address:         req.Address,
		Description:     req.Description,
		FinedBy:         req.UserId,
	}

	err := s.repo.CreateFineHistory(ctx, payload)
	s.handleErrorExitHistory(err, false)
	s.exception.PanicIfError(err, false)
}

func (s *historyService) GetAllHistoryStatistic(ctx context.Context, req request.TimeFrequency) response.HistoryStatistic {
	timeFrequency := req.BuildTimeFrequency()

	idx := 0
	var tempArr = make(map[string]int)
	var res response.HistoryStatistic
	for date := timeFrequency.StartDate; !date.After(timeFrequency.EndDate); date = timeFrequency.CallbackDate(date) {
		name := timeFrequency.CallbackName(date)
		if _, ok := tempArr[name]; !ok {
			tempArr[name] = idx
		}
		res.Charts = append(res.Charts, response.HistoryStatisticChart{
			Name:    name,
			Revenue: 0,
			Vehicle: 0,
		})
		idx++
	}

	resStatistic := sql.GetCountHistoryStatisticParams{
		StartAt: utils.PGTimeStamp(timeFrequency.StartDate),
		EndAt:   utils.PGTimeStamp(timeFrequency.EndDate),
	}
	calc, err := s.repo.GetCountHistoryStatistic(ctx, resStatistic)
	s.exception.PanicIfError(err, true)
	setting := s.serviceSetting.GetAllSetting(ctx)

	res.RevenueTotal = calc.Revenue
	res.VehicleTotal = int(calc.Total)
	res.CurrentVehicle = int(calc.EntryTotal)
	res.ExitRevenue = calc.ExitRevenue
	res.ExitTotal = int(calc.ExitRevenue)
	res.FineRevenue = calc.FineRevenue
	res.FineTotal = int(calc.FineTotal)

	if req.TimeFrequency == constants.FilterTimeFrequencyToday {
		res.AvailableSpace = setting.MaxCapacity - int(calc.EntryTotal)
	}

	payload := model.AllHistoryStatistic{
		StartDate:     timeFrequency.StartDate,
		EndDate:       timeFrequency.EndDate,
		TimeFrequency: req.TimeFrequency,
	}
	data, err := s.repo.AllHistoryStatistic(ctx, payload)
	s.exception.PanicIfError(err, true)

	for _, data := range data {
		if !data.Date.Valid {
			continue
		}

		name := timeFrequency.CallbackName(data.Date.Time)
		if idx, ok := tempArr[name]; ok {
			res.Charts[idx].Revenue += float64(data.Revenue)
			res.Charts[idx].Vehicle += int(data.Vehicle)
		}
	}

	return res
}
