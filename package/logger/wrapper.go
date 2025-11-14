package logger

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

type wrapper struct {
	zerolog.Logger
	prefix string
}

func (w *wrapper) handleMessage(format string, args []any) string {
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

func (w *wrapper) Error(err error, args ...any) {
	w.Logger.Error().Err(err).Msg(w.handleMessage("", args))
}

func (w *wrapper) Errorf(err error, format string, args ...any) {
	w.Logger.Error().Err(err).Msg(w.handleMessage(format, args))
}

func (w *wrapper) Fatal(err error, args ...any) {
	w.Logger.Fatal().Err(err).Msg(w.handleMessage("", args))
}

func (w *wrapper) Fatalf(err error, format string, args ...any) {
	w.Logger.Fatal().Err(err).Msg(w.handleMessage(format, args))
}

func (w *wrapper) Info(args ...any) {
	w.Logger.Info().Msg(w.handleMessage("", args))
}

func (w *wrapper) Infof(format string, args ...any) {
	w.Logger.Info().Msg(w.handleMessage(format, args))
}

func (w *wrapper) Debug(args ...any) {
	w.Logger.Debug().Msg(w.handleMessage("", args))
}

func (w *wrapper) Debugf(format string, args ...any) {
	w.Logger.Debug().Msg(w.handleMessage(format, args))
}

func (w *wrapper) Warn(args ...any) {
	w.Logger.Warn().Msg(w.handleMessage("", args))
}

func (w *wrapper) Warnf(format string, args ...any) {
	w.Logger.Warn().Msg(w.handleMessage(format, args))
}
