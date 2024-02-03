package lint

import (
	"os"

	"github.com/rinchsan/groupvar"
	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

type Linter struct{}

func NewLinter() *Linter {
	return &Linter{}
}

var analyzers = []*analysis.Analyzer{
	groupvar.Analyzer,
}

func (l *Linter) Run() error {
	os.Args = slices.Delete(os.Args, 1, 2)
	multichecker.Main(analyzers...)
	return nil
}
