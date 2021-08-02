package codecheck_test

import (
	"testing"

	"github.com/uudashr/go-playground/codecheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	codecheck.Analyzer.Flags.Set("limit", "0")
	analysistest.Run(t, testdata, codecheck.Analyzer, "a")
}

func TestAnalyzerLimit2(t *testing.T) {
	testdata := analysistest.TestData()
	codecheck.Analyzer.Flags.Set("limit", "2")
	analysistest.Run(t, testdata, codecheck.Analyzer, "b")
}

func TestAnalyzerDefaultLimit(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, codecheck.Analyzer, "b")
}
