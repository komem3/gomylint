package main

import (
	"github.com/komem3/ctxfirst"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(ctxfirst.Analyzer)
}
