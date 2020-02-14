package linter

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

type isWrapper struct{}

func (f *isWrapper) AFact() {}

var Analyzer = &analysis.Analyzer{
	Name:     "flextimelinter",
	Doc:      Doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
}

const Doc = "flextimelinter is ..."

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		foundTimeImport, foundFlexTime := findImports(pass, f)
		foundNotCallImport := false
		ast.Inspect(f, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.SelectorExpr:
				if pkgIdent, ok := n.X.(*ast.Ident); ok {
					pkgName, ok := pass.TypesInfo.Uses[pkgIdent].(*types.PkgName)
					if !ok || pkgName.Imported().Path() != "time" {
						return false
					}
					foundNotCallImport = true
					return false
				}
			case *ast.CallExpr:
				if s, ok := n.Fun.(*ast.SelectorExpr); ok {
					if pkgIdent, ok := s.X.(*ast.Ident); ok {
						pkgName, ok := pass.TypesInfo.Uses[pkgIdent].(*types.PkgName)
						if !ok || pkgName.Imported().Path() != "time" {
							break
						}

						d := analysis.Diagnostic{
							Pos:     n.Pos(),
							Message: "Prefer use flextime",
							SuggestedFixes: []analysis.SuggestedFix{
								{
									Message: "Prefer use flextime",
									TextEdits: []analysis.TextEdit{
										{
											Pos:     pkgIdent.Pos(),
											End:     pkgIdent.End(),
											NewText: []byte(getFlexTimeName(foundFlexTime)),
										},
									},
								},
							},
						}
						pass.Report(d)
						return false
					}
				}
			}
			return true
		})
		if !foundNotCallImport {
			if foundTimeImport != nil && foundFlexTime == nil {
				d := analysis.Diagnostic{
					Pos:     foundTimeImport.Pos(),
					Message: "Prefer use flextime",
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "Prefer use flextime",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     foundTimeImport.Pos(),
									End:     foundTimeImport.End(),
									NewText: []byte(`"github.com/Songmu/flextime"`),
								},
							},
						},
					},
				}
				pass.Report(d)
			}
		}
	}
	return nil, nil
}

func getFlexTimeName(pkg *types.PkgName) string {
	if pkg != nil {
		return pkg.Name()
	}
	return "github.com/Songmu/flextime"
}

func findImports(pass *analysis.Pass, f *ast.File) (*ast.ImportSpec, *types.PkgName) {
	var foundTimeImport *ast.ImportSpec = nil
	var foundFlexTime *types.PkgName = nil
	for _, e := range f.Imports {
		pkg, ok := pass.TypesInfo.Defs[e.Name]
		if ok {
			if pkg.(*types.PkgName).Imported().Path() == "time" {
				foundTimeImport = e
			}
			if pkg.(*types.PkgName).Imported().Path() == "github.com/Songmu/flextime" {
				foundFlexTime = pkg.(*types.PkgName)
			}
		} else if e.Path.Value == `"time"` {
			foundTimeImport = e
		}
	}
	return foundTimeImport, foundFlexTime
}
