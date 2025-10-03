package allowimports

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// TODO Turn this allow list into config or command line input
var (
	allowPrefixes = []string{"github.com/Stacky32", "bytes"}
)

var Analyzer = &analysis.Analyzer{
	Name:     "allowimports",
	Doc:      "Validates imports against a whitelist",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(p *analysis.Pass) (any, error) {
	filter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	inspect := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(filter, func(node ast.Node) {
		imp, ok := node.(*ast.ImportSpec)
		if !ok {
			pos := p.Fset.Position(node.Pos())
			fmt.Fprintf(os.Stderr, "%s:%d: warning: not an import\n", pos.Filename, pos.Line)
			return
		}

		path := strings.Trim(imp.Path.Value, `"`)
		if isAllowed(path) {
			return
		}

		p.Reportf(imp.Pos(), "importing forbidden package %q", path)
	})

	return nil, nil
}

func isAllowed(path string) bool {
	for _, p := range allowPrefixes {
		if strings.HasPrefix(path, p) {
			return true
		}
	}

	return false
}
