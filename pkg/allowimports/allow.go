package allowimports

import (
	"errors"
	"fmt"
	"go/ast"
	"os"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Allow []string `yaml:"allow"`
}

var (
	configPath    string
	allowPrefixes []string
)

var Analyzer = &analysis.Analyzer{
	Name:     "allowimports",
	Doc:      "Validates imports against a whitelist",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", "", "path to YAML config file (required)")
}

func run(p *analysis.Pass) (any, error) {
	if configPath == "" {
		return nil, errors.New("missing flag `-config` - YAML config required")
	}

	if len(allowPrefixes) == 0 {
		cfg, err := loadConfig(configPath)
		if err != nil {
			return nil, err
		}
		allowPrefixes = cfg.Allow
	}

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
		for _, p := range allowPrefixes {
			if strings.HasPrefix(path, p) {
				return
			}
		}

		p.Reportf(imp.Pos(), "importing forbidden package %q", path)
	})

	return nil, nil
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	return &cfg, nil
}
