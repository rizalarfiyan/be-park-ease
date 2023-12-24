package exception

import (
	"be-park-ease/internal/response"
	"be-park-ease/logger"
	"be-park-ease/utils"
	"errors"
	"fmt"
	"net/http"
	"runtime"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

type Exception interface {
	DefaultData(isList bool) interface{}
	PanicIfError(err error, isList bool)
	PanicIfErrorWithoutNoSqlResult(err error, isList bool)
	IsNotFoundMessage(value interface{}, message string, isList bool)
	IsNotFound(value interface{}, isList bool, modules ...string)
	IsUnprocessableEntity(value interface{}, message string, isList bool)
	IsBadRequest(value interface{}, message string, isList bool)
	IsBadRequestErr(err error, message string, isList bool)
	IsUnauthorized(message string, isList bool)
	IsForbidden(message string, isList bool)
}

type exception struct {
	log *zerolog.Logger
}

func NewException(types string) Exception {
	return &exception{
		log: logger.GetWithoutCaller(types),
	}
}

func (e *exception) DefaultData(isList bool) interface{} {
	if isList {
		return []any{}
	}
	return nil
}

func (e *exception) getCaller(skips ...int) string {
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)

}

func (e *exception) PanicIfError(err error, isList bool) {
	if err != nil {
		data := e.DefaultData(isList)
		e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, err.Error(), data))
	}
}

func (e *exception) PanicIfErrorWithoutNoSqlResult(err error, isList bool) {
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		data := e.DefaultData(isList)
		e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, err.Error(), data))
	}
}

func (e *exception) isEmptyCallback(value interface{}, isList bool, callback func(isList bool, data interface{}) *response.BaseResponse) {
	if utils.IsEmptyAll(value) {
		data := e.DefaultData(isList)
		panic(callback(isList, data))
	}
}

func (e *exception) IsNotFoundMessage(value interface{}, message string, isList bool) {
	e.isEmptyCallback(value, isList, func(isList bool, data interface{}) *response.BaseResponse {
		if isList {
			return response.NewErrorMessage(http.StatusOK, message, data)
		}
		return response.NewErrorMessage(http.StatusNotFound, message, data)
	})
}

func (e *exception) IsNotFound(value interface{}, isList bool, modules ...string) {
	module := "Data"
	if len(modules) > 0 {
		module = modules[0]
	}
	e.IsNotFoundMessage(value, fmt.Sprintf("%s not found", module), isList)
}

func (e *exception) IsUnprocessableEntity(value interface{}, message string, isList bool) {
	e.isEmptyCallback(value, isList, func(isList bool, data interface{}) *response.BaseResponse {
		return response.NewErrorMessage(http.StatusUnprocessableEntity, message, data)
	})
}

func (e *exception) IsBadRequest(value interface{}, message string, isList bool) {
	e.isEmptyCallback(value, isList, func(isList bool, data interface{}) *response.BaseResponse {
		return response.NewErrorMessage(http.StatusBadRequest, message, data)
	})
}

func (e *exception) IsBadRequestErr(err error, message string, isList bool) {
	if err != nil {
		e.IsBadRequest(nil, message, isList)
	}
}

func (e *exception) IsUnauthorized(message string, isList bool) {
	data := e.DefaultData(isList)
	panic(response.NewErrorMessage(http.StatusInternalServerError, message, data))
}

func (e *exception) IsForbidden(message string, isList bool) {
	data := e.DefaultData(isList)
	panic(response.NewErrorMessage(http.StatusForbidden, message, data))
}
