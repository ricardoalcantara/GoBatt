package logger

import (
	"path/filepath"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() ILogger {
	logger := Logger{}
	logger.Init()
	return &logger
}

type ILogger interface {
	Init()
	Info(message string)
	Error(message string)
}

type Logger struct {
	logger zerolog.Logger
}

func (l *Logger) Init() {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}
	l.logger = log.With().Caller().Logger()
}

func (l *Logger) Info(message string) {
	log.Info().Msg(message)
}

func (l *Logger) Error(message string) {
	log.Error().Msg(message)
}
