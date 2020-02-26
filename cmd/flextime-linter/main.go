package main

import (
	linter "github.com/wreulicke/flextime-linter"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(linter.Analyzer) }
