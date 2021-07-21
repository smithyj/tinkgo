package logx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	logger *zap.Logger
)

func Setup(c Config) {
	if c.Path == "" {
		c.Path = "./logs"
	}
	if c.Name == "" {
		c.Name = "app"
	}
	hook := lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", c.Path, c.Name),
		MaxSize:    c.MaxSize,
		MaxAge:     c.MaxAge,
		MaxBackups: c.MaxBackups,
		Compress:   c.Compress,
	}

	syncer := zapcore.AddSync(&hook)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	level := zap.InfoLevel
	if c.Level != "" {
		_ = level.UnmarshalText([]byte(c.Level))
	}
	var cores []zapcore.Core
	var ws []zapcore.WriteSyncer
	ws = append(ws, zapcore.AddSync(syncer))
	if c.Debug {
		level = zap.DebugLevel
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cores = append(cores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		))
	} else {
		// 非 DEBUG 模式，输出 JSON
		ws = append(ws, zapcore.AddSync(os.Stdout))
	}

	cores = append(cores, zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(ws...),
		level,
	))

	logger = zap.New(
		zapcore.NewTee(cores...),
		zap.AddStacktrace(zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zap.ErrorLevel
		})),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
}

func Sync() {
	defer func() {
		_ = logger.Sync()
	}()
}