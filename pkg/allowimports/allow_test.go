package allowimports_test

import (
	"testing"

	"github.com/Stacky32/go-analyser/pkg/allowimports"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	td := analysistest.TestData()
	t.Logf("Blah: %s\n", td)
	analysistest.Run(t, td, allowimports.Analyzer)
}
