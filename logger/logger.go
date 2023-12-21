package logger

import (
	"be-park-ease/config"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

var once sync.Once
var log zerolog.Logger

func Init(conf *config.Config) {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339

		var output io.Writer = zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
		}

		if conf.Logger.IsLogRotator {
			lumberjakLog := NewLumberjackLogger(conf)
			output = zerolog.MultiLevelWriter(output, lumberjakLog.Run())
		}

		log = zerolog.New(output).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	})
}

func UpdateLogLevel(level zerolog.Level) {
	log = log.Level(level)
}

func Get(types string) zerolog.Logger {
	return log.With().Str("type", types).Caller().Logger()
}

func GetWithoutCaller(types string) *zerolog.Logger {
	logg := log.With().Str("type", types).Logger()
	return &logg
}
