package main

import (
	"github.com/komem3/alphabetorder"
	"github.com/komem3/cmpcheck"
	"github.com/komem3/ctxfirst"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		ctxfirst.Analyzer,
		cmpcheck.Analyzer,
		alphabetorder.Analyzer,
	)
}
