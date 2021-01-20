package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path/filepath"
)

type Logger struct {
	zerolog.Logger
}

const (
	moduleFieldName = "module"
)

// ConfigureLog creates new logger based on github.com/rs/zerolog package
func ConfigureLog(logFile, logLevel, module string) (*Logger, error) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.DebugLevel
	}

	logWriter, err := getLogWriter(logFile)
	if err != nil {
		return nil, err
	}

	logger := zerolog.New(
		zerolog.ConsoleWriter{
			Out:        logWriter,
			NoColor:    true,
			TimeFormat: "2006-01-02 15:04:05.000",
			PartsOrder: []string{zerolog.TimestampFieldName, moduleFieldName, zerolog.LevelFieldName, zerolog.MessageFieldName},
		},
	).Level(level).With().Str(moduleFieldName, module).Logger()
	return &Logger{logger}, nil
}

func getLogWriter(logFileName string) (io.Writer, error) {
	if logFileName == "stdout" || logFileName == "" {
		return os.Stdout, nil
	}

	logDir := filepath.Dir(logFileName)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("can't create log directories %s: %s", logDir, err.Error())
	}
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("can't open log file %s: %s", logFileName, err.Error())
	}
	return logFile, nil
}

func (l Logger) Debug(args ...interface{}) {
	l.Logger.Debug().Timestamp().Msg(fmt.Sprint(args))
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Timestamp().Msgf(format, args)
}

func (l Logger) Info(args ...interface{}) {
	l.Logger.Info().Timestamp().Msg(fmt.Sprint(args))
}

func (l Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info().Timestamp().Msgf(format, args)
}

func (l Logger) Error(args ...interface{}) {
	l.Logger.Error().Timestamp().Msgf(fmt.Sprint(args))
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Timestamp().Msgf(format, args)
}

func (l Logger) Fatal(args ...interface{}) {
	l.Logger.Fatal().Timestamp().Msg(fmt.Sprint(args))
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal().Timestamp().Msgf(format, args)
}

func (l Logger) Warning(args ...interface{}) {
	l.Logger.Warn().Timestamp().Msg(fmt.Sprint(args))
}

func (l Logger) Warningf(format string, args ...interface{}) {
	l.Logger.Warn().Timestamp().Msgf(format, args)
}
