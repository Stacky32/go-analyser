//go:build plugin
// +build plugin

package main

import (
	"github.com/Stacky32/go-analyser/pkg/allowimports"
	"golang.org/x/tools/go/analysis"
)

var AnalyzerPlugin = map[string]*analysis.Analyzer{
	"allowimports": allowimports.Analyzer,
}
