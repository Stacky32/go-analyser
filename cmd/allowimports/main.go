package main

import (
	"github.com/Stacky32/go-analyser/pkg/allowimports"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(allowimports.Analyzer)
}
