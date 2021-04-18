package main

import (
	"github.com/komem3/alphabetical"
	"github.com/komem3/cmpcheck"
	"github.com/komem3/ctxfirst"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		ctxfirst.Analyzer,
		cmpcheck.Analyzer,
		alphabetical.Analyzer,
	)
}
