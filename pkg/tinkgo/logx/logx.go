package logx

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var event = map[zerolog.Level]*zerolog.Event {}

func Debug() *zerolog.Event {
	return event[zerolog.DebugLevel]
}

func Info() *zerolog.Event {
	return event[zerolog.InfoLevel]
}

func Warn() *zerolog.Event {
	return event[zerolog.WarnLevel]
}

func Error() *zerolog.Event {
	return event[zerolog.ErrorLevel]
}

func Fatal() *zerolog.Event {
	return event[zerolog.FatalLevel]
}

func Panic() *zerolog.Event {
	return event[zerolog.PanicLevel]
}

func Trace() *zerolog.Event {
	return event[zerolog.TraceLevel]
}

func Setup(config Config) {
	if config.Path == "" {
		config.Path = "./logs"
	}
	level := map[zerolog.Level]func(logger zerolog.Logger)*zerolog.Event{
		zerolog.DebugLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Debug()
		},
		zerolog.InfoLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Info()
		},
		zerolog.WarnLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Warn()
		},
		zerolog.ErrorLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Error()
		},
		zerolog.FatalLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Fatal()
		},
		zerolog.PanicLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Panic()
		},
		zerolog.TraceLevel: func(logger zerolog.Logger) *zerolog.Event {
			return logger.Trace()
		},
	}

	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	for k, l := range level {
		hook := lumberjack.Logger{
			Filename:   fmt.Sprintf("%s/%s.log", config.Path, zerolog.LevelFieldMarshalFunc(k)),
			MaxSize:    config.MaxSize,
			MaxAge:     config.MaxAge,
			MaxBackups: config.MaxBackups,
			Compress:   config.Compress,
		}
		writer := zerolog.MultiLevelWriter(
			os.Stdout,
			&hook,
		)
		logger := zerolog.New(writer)
		event[k] = l(logger)
	}
}