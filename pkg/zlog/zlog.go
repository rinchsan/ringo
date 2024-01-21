package zlog

import (
	"io"
	"runtime"

	"golang.org/x/exp/slog"
)

type Logger struct {
	l *slog.Logger
}

func NewLogger(w io.Writer) *Logger {
	h := slog.NewJSONHandler(w, nil)
	l := slog.New(h)
	return &Logger{
		l: l,
	}
}

func (l *Logger) Info(message string, args ...any) {
	pc, file, line, _ := runtime.Caller(1)
	args = append(args, "call", runtime.FuncForPC(pc).Name(), "file", file, "line", line)
	l.l.Info(message, args...)
}

func (l *Logger) Error(err error, args ...any) {
	pc, file, line, _ := runtime.Caller(1)
	args = append(args, "call", runtime.FuncForPC(pc).Name(), "file", file, "line", line)
	l.l.Error(err.Error(), args...)
}
