package zlog_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/rinchsan/ringo/pkg/zlog"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestLogger_Info(t *testing.T) {
	t.Parallel()

	buf := bytes.NewBuffer(nil)
	l := zlog.NewLogger(buf)

	attrs := map[string]any{
		"msg":  "test info",
		"call": "github.com/rinchsan/ringo/pkg/zlog_test.TestLogger_Info",
		"file": "/pkg/zlog/zlog_test.go",
		"line": float64(38),
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
	_ = json.Unmarshal(buf.Bytes(), &got)

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
	t.Parallel()

	buf := bytes.NewBuffer(nil)
	l := zlog.NewLogger(buf)

	err := errors.New("test error")
	attrs := map[string]any{
		"msg":  "test error",
		"call": "github.com/rinchsan/ringo/pkg/zlog_test.TestLogger_Error",
		"file": "/pkg/zlog/zlog_test.go",
		"line": float64(84),
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
	_ = json.Unmarshal(buf.Bytes(), &got)

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

func TestRepoBase(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		filepath string
		want     string
	}{
		"filepath on local": {
			filepath: "/Users/rinchsan/go/src/github.com/rinchsan/ringo/pkg/zlog/zlog_test.go",
			want:     "/pkg/zlog/zlog_test.go",
		},
		"filepath on GitHub Actions": {
			filepath: "/home/runner/work/ringo/ringo/pkg/zlog/zlog_test.go",
			want:     "/pkg/zlog/zlog_test.go",
		},
	}

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := zlog.RepoBase(c.filepath)
			if got != c.want {
				t.Fatalf("want %#v, but got %#v", c.want, got)
			}
		})
	}
}
