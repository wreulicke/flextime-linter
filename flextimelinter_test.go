package linter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPrint(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "testdata/src/a/a.go", nil, parser.Mode(0))
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range f.Decls {
		ast.Print(fset, d)
	}
	t.Error("test")
}

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "a")
	t.Error("test")
}
