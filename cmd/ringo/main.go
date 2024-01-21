package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/rinchsan/ringo/pkg/rest"
	"github.com/rinchsan/ringo/pkg/zlog"
	_ "go.uber.org/automaxprocs"
)

type ringo struct {
	REST CmdREST `cmd:"rest" help:"Run REST Server"`
}

type CmdREST struct{}

func (c *CmdREST) Run() error {
	logger := zlog.NewLogger(os.Stdout)
	return rest.NewServer(logger).Run()
}

func main() {
	var ringo ringo
	ctx := kong.Parse(&ringo)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
