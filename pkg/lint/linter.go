package lint

import (
	"os"

	"github.com/gordonklaus/ineffassign/pkg/ineffassign"
	"github.com/kisielk/errcheck/errcheck"
	"github.com/rinchsan/groupvar"
	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/simple"
)

type Linter struct {
	analyzers []*analysis.Analyzer
}

func NewLinter() *Linter {
	analyzers := []*analysis.Analyzer{
		groupvar.Analyzer,
		errcheck.Analyzer,
		ineffassign.Analyzer,
	}

	for _, v := range simple.Analyzers {
		analyzers = append(analyzers, v.Analyzer)
	}

	return &Linter{
		analyzers: analyzers,
	}
}

func (l *Linter) Run() error {
	os.Args = slices.Delete(os.Args, 1, 2)
	multichecker.Main(l.analyzers...)
	return nil
}
