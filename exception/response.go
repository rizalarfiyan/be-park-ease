package exception

import (
	"be-park-ease/internal/response"
	"be-park-ease/logger"
	"fmt"
	"net/http"
	"runtime"

	"github.com/rs/zerolog"
)

type Exception interface {
	PanicIfError(err error, isList bool)
}

type exception struct {
	log *zerolog.Logger
}

func NewException(types string) Exception {
	return &exception{
		log: logger.GetWithoutCaller(types),
	}
}

func (e *exception) defaultData(isList bool) interface{} {
	if isList {
		return []any{}
	}
	return nil
}

func (e *exception) getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d\n", file, line)

}

func (e *exception) PanicIfError(err error, isList bool) {
	if err != nil {
		data := e.defaultData(isList)
		e.log.Error().Str("caller", e.getCaller()).Err(err).Msg("SERVER ERROR")
		panic(response.NewErrorMessage(http.StatusInternalServerError, err.Error(), data))
	}
}
