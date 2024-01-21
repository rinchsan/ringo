package zlog_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/rinchsan/ringo/pkg/zlog"
)

func TestLogger_Info(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := zlog.NewLogger(buf)

	attrs := map[string]any{
		"msg":  "test info",
		"call": "github.com/rinchsan/ringo/pkg/zlog_test.TestLogger_Info",
		"file": "/pkg/zlog/zlog_test.go",
		"line": float64(31),
	}
	argsMap := map[string]any{
		"foo": "bar",
		"baz": float64(123),
	}
	var args []any
	for k, v := range argsMap {
		args = append(args, k, v)
	}
	l.Info(attrs["msg"].(string), args...)

	got := make(map[string]any)
	json.Unmarshal(buf.Bytes(), &got)

	now, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", got["time"].(string))
	if err != nil {
		t.Fatal(err)
	}
	if time.Now().Sub(now) > time.Second {
		t.Fatal("'time' does not describe current time")
	}

	for k := range attrs {
		if got[k] != attrs[k] {
			t.Fatalf("want %#v, but got %#v for key '%s'", attrs[k], got[k], k)
		}
	}
	for k := range argsMap {
		if got[k] != argsMap[k] {
			t.Fatalf("want %#v, but got %#v for key '%s'", argsMap[k], got[k], k)
		}
	}
}

func TestLogger_Error(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	l := zlog.NewLogger(buf)

	err := errors.New("test error")
	attrs := map[string]any{
		"msg":  "test error",
		"call": "github.com/rinchsan/ringo/pkg/zlog_test.TestLogger_Error",
		"file": "/pkg/zlog/zlog_test.go",
		"line": float64(75),
	}
	argsMap := map[string]any{
		"foo": "bar",
		"baz": float64(123),
	}
	var args []any
	for k, v := range argsMap {
		args = append(args, k, v)
	}
	l.Error(err, args...)

	got := make(map[string]any)
	json.Unmarshal(buf.Bytes(), &got)

	now, err := time.Parse("2006-01-02T15:04:05.999999Z07:00", got["time"].(string))
	if err != nil {
		t.Fatal(err)
	}
	if time.Now().Sub(now) > time.Second {
		t.Fatal("'time' does not describe current time")
	}

	for k := range attrs {
		if got[k] != attrs[k] {
			t.Fatalf("want %#v, but got %#v for key '%s'", attrs[k], got[k], k)
		}
	}
	for k := range argsMap {
		if got[k] != argsMap[k] {
			t.Fatalf("want %#v, but got %#v for key '%s'", argsMap[k], got[k], k)
		}
	}
}
