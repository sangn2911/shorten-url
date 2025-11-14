package logger

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
	prefix string
}

func (w *Logger) handleMessage(format string, args []any) string {
	var builder strings.Builder
	if len(w.prefix) > 0 {
		builder.WriteString(w.prefix)
		builder.WriteString(" ")
	}
	if len(format) > 0 {
		builder.WriteString(fmt.Sprintf(format, args...))
	} else {
		builder.WriteString(fmt.Sprint(args...))
	}
	return builder.String()
}

func (w *Logger) Error(err error, args ...any) {
	w.Logger.Error().Err(err).Msg(w.handleMessage("", args))
}

func (w *Logger) Errorf(err error, format string, args ...any) {
	w.Logger.Error().Err(err).Msg(w.handleMessage(format, args))
}

func (w *Logger) Fatal(err error, args ...any) {
	w.Logger.Fatal().Err(err).Msg(w.handleMessage("", args))
}

func (w *Logger) Fatalf(err error, format string, args ...any) {
	w.Logger.Fatal().Err(err).Msg(w.handleMessage(format, args))
}

func (w *Logger) Info(args ...any) {
	w.Logger.Info().Msg(w.handleMessage("", args))
}

func (w *Logger) Infof(format string, args ...any) {
	w.Logger.Info().Msg(w.handleMessage(format, args))
}

func (w *Logger) Debug(args ...any) {
	w.Logger.Debug().Msg(w.handleMessage("", args))
}

func (w *Logger) Debugf(format string, args ...any) {
	w.Logger.Debug().Msg(w.handleMessage(format, args))
}

func (w *Logger) Warn(args ...any) {
	w.Logger.Warn().Msg(w.handleMessage("", args))
}

func (w *Logger) Warnf(format string, args ...any) {
	w.Logger.Warn().Msg(w.handleMessage(format, args))
}

func (w *Logger) WithContext(ctx context.Context) *Logger {
	pid := GetPID(ctx)
	if len(pid) == 0 {
		return w
	}
	l := *w
	l.prefix = fmt.Sprintf("request_id=%s", pid)
	return &l
}
