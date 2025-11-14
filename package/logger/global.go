package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var _logger wrapper

func init() {
	InitZeroLogger(zerolog.InfoLevel)
}

func InitZeroLogger(level zerolog.Level) {
	writer := zerolog.ConsoleWriter{
		Out:              os.Stdout,
		TimeFormat:       time.RFC3339,
		NoColor:          false,
		FormatTimestamp:  nil,
		FormatLevel:      nil,
		FormatFieldName:  nil,
		FormatFieldValue: nil,
		FormatMessage: func(i any) string {
			return fmt.Sprintf("%v", i)
		},
	}
	_logger = wrapper{
		Logger: zerolog.New(writer).With().Timestamp().Logger().Level(level),
	}
}

func Error(err error, args ...any) {
	_logger.Error(err, args...)
}

func Errorf(err error, format string, args ...any) {
	_logger.Errorf(err, format, args)
}

func Fatal(err error, args ...any) {
	_logger.Fatal(err, args...)
}

func Fatalf(err error, format string, args ...any) {
	_logger.Fatalf(err, format, args...)
}

func Info(args ...any) {
	_logger.Info(args...)
}

func Infof(format string, args ...any) {
	_logger.Infof(format, args...)
}

func Debug(args ...any) {
	_logger.Debug(args...)
}

func Debugf(format string, args ...any) {
	_logger.Debugf(format, args...)
}

func Warn(args ...any) {
	_logger.Warn(args...)
}

func Warnf(format string, args ...any) {
	_logger.Warnf(format, args...)
}

func JSONFormat(value any) string {
	prefix := ""
	indent := "   "
	data, err := json.MarshalIndent(value, prefix, indent)
	if err != nil {
		Errorf(err, "fail to parse %v", value)
		return ""
	}
	return string(data)
}
