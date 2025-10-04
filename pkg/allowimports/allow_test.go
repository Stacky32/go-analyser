package allowimports_test

import (
	"path/filepath"
	"testing"

	"github.com/Stacky32/go-analyser/pkg/allowimports"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	td := analysistest.TestData()
	configPath := filepath.Join(td, "allow.yaml")
	if err := allowimports.Analyzer.Flags.Set("config", configPath); err != nil {
		t.Fatalf("Failed to set flag `-config`: %+v", err)
	}

	analysistest.Run(t, td, allowimports.Analyzer)
}
