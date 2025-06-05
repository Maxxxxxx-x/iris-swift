package logger

import (
	"io"
	"os"
	"time"

	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Log() *zerolog.Event
	Fatal() *zerolog.Event
	Err(err error) *zerolog.Event
	Panic() *zerolog.Event
	Error() *zerolog.Event
	Warn() *zerolog.Event
	Info() *zerolog.Event
	Trace() *zerolog.Event
	Debug() *zerolog.Event
	With() zerolog.Context
	SetLogLevel(level string)
	Printf(format string, v ...interface{})
	Print(v ...interface{})
}

type DefaultLogger struct {
	log     zerolog.Logger
	level   zerolog.Level
	writers []io.Writer
}

var logger *DefaultLogger

func (logger *DefaultLogger) SetLogLevel(level string) {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(lvl)
}

func Log() *zerolog.Event {
	return getDefaultLogger().log.Log()
}

func Fatal() *zerolog.Event {
	return getDefaultLogger().log.Fatal()
}

func Error() *zerolog.Event {
	return getDefaultLogger().log.Error()
}

func Err(err error) *zerolog.Event {
	return getDefaultLogger().log.Err(err)
}

func Warn() *zerolog.Event {
	return getDefaultLogger().log.Warn()
}

func Info() *zerolog.Event {
	return getDefaultLogger().log.Info()
}

func Debug() *zerolog.Event {
	return getDefaultLogger().log.Debug()
}

func Panic() *zerolog.Event {
	return getDefaultLogger().log.Panic()
}

func Trace() *zerolog.Event {
	return getDefaultLogger().log.Trace()
}

func Print(v ...interface{}) {
	getDefaultLogger().log.Print(v...)
}

func Printf(format string, v ...interface{}) {
	getDefaultLogger().log.Printf(format, v...)
}

func With() zerolog.Context {
	return getDefaultLogger().log.With().Caller()
}

func getDefaultLogger() *DefaultLogger {
	if logger == nil {
		panic("Logger: package not initialized. Please initialize the logger first!")
	}
	return logger
}

func NewLogger(serviceName string) zerolog.Logger {
	return getDefaultLogger().log.With().Str("service", serviceName).Logger()
}

func Init(cfg config.LoggerConfig, env string) {
	logger = &DefaultLogger{
		writers: make([]io.Writer, 0),
	}

	lvl, err := zerolog.ParseLevel(cfg.Log_Level)
	if err != nil {
		lvl = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(lvl)

	if env == "dev" {
		logger.writers = append(logger.writers, zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
		})
	} else {
		logger.writers = append(logger.writers, os.Stderr)
	}

	if cfg.Log_File_Path != "" {
		logger.writers = append(logger.writers, &lumberjack.Logger{
			Filename:   cfg.Log_File_Path,
			Compress:   cfg.Compress_Logs,
			MaxBackups: cfg.Max_Backups,
			MaxSize:    cfg.Max_Size,
			MaxAge:     cfg.Max_Age,
		})
	}

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	logger.log = zerolog.New(io.MultiWriter(logger.writers...)).
		With().
		Timestamp().
		Stack().
		Logger()
}
