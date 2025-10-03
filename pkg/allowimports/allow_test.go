package allowimports_test

import (
	"testing"

	"github.com/Stacky32/go-analyser/pkg/allowimports"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), allowimports.Analyzer)
}
