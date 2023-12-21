package logger

import (
	"be-park-ease/config"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

type LumberjackLog interface {
	Run() *lumberjack.Logger
}

type lumberjackLog struct {
	LogPath       string
	CompressLog   bool
	DailyRotate   bool
	SleepDuration time.Duration
	lastLogDate   string
	lumberjackLog *lumberjack.Logger
}

func NewLumberjackLogger(conf *config.Config) LumberjackLog {
	return &lumberjackLog{
		LogPath:       conf.Logger.Path,
		DailyRotate:   conf.Logger.IsDailyRotate,
		CompressLog:   conf.Logger.IsCompressed,
		SleepDuration: conf.Logger.SleepDuration,
	}
}

func (l *lumberjackLog) Run() *lumberjack.Logger {
	l.lumberjackLog = &lumberjack.Logger{
		Filename:  l.LogPath,
		Compress:  l.CompressLog,
		LocalTime: true,
	}

	l.lastLogDate = time.Now().Format(time.DateOnly)

	if l.DailyRotate {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			l.rotate()
		}()
	}

	return l.lumberjackLog
}

func (l *lumberjackLog) rotate() {
	ticker := time.NewTicker(l.SleepDuration)
	defer ticker.Stop()

	for range ticker.C {
		if l.lumberjackLog == nil {
			continue
		}

		now := time.Now().Format(time.DateOnly)
		if l.lastLogDate != now {
			l.lastLogDate = now
			l.lumberjackLog.Rotate()
		}
	}
}
