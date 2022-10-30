package main

import (
	"github.com/alecthomas/kong"
	"github.com/rinchsan/ringo/pkg/rest"
	_ "go.uber.org/automaxprocs"
)

type ringo struct {
	REST CmdREST `cmd:"rest" help:"Run REST Server"`
}

type CmdREST struct{}

func (c *CmdREST) Run() error {
	return rest.NewServer().Run()
}

func main() {
	var ringo ringo
	ctx := kong.Parse(&ringo)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
