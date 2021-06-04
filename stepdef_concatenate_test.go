package main

import "github.com/cucumber/godog"

func aPDFFileWithPages(arg1 int) error {
	return godog.ErrPending
}

func aSecondPDFFileWithPages(arg1 int) error {
	return godog.ErrPending
}

func pDFFilesAreConcatenated() error {
	return godog.ErrPending
}

func theResultingFileHasFourPages() error {
	return godog.ErrPending
}

func InitializeConcatenateScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^A PDF file with (\d+) pages$`, aPDFFileWithPages)
	ctx.Step(`^a second PDF file with (\d+) pages$`, aSecondPDFFileWithPages)
	ctx.Step(`^PDF files are concatenated$`, pDFFilesAreConcatenated)
	ctx.Step(`^the resulting file has four pages$`, theResultingFileHasFourPages)
}
