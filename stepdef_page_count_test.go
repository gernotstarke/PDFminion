package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"pdfminion/fileutil"
)

var pageCount int
var pdfFileName string

func aDirUnderSampleContainingFile(directory, samplePDFDir, pdfFile string) error {
	pdfFileName = samplePDFDir + "/" + directory + "/" + pdfFile

	fileExists, err := fileutil.FileExists(samplePDFDir)

	if (err != nil) || !fileExists {
		return fmt.Errorf("sample PDF file does not exist" + err.Error())
	} else {
		return nil
	}
}

func numberOfPagesIsCounted() error {
	pageCount, _ = fileutil.CountPagesOfPDFFile(pdfFileName)

	return nil
}

func theNumberOfPagesShouldBe(expectedPageCount int) error {
	if expectedPageCount != pageCount {
		return fmt.Errorf("expected %d nr of pages, but found %d", expectedPageCount, pageCount)
	} else {
		return nil
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		pageCount = -1
	})

	ctx.Step(`^A "([^"]*)" under "([^"]*)" containing "([^"]*)"$`, aDirUnderSampleContainingFile)
	ctx.Step(`^Number of pages is counted$`, numberOfPagesIsCounted)
	ctx.Step(`^the number of pages should be (\d+)$`, theNumberOfPagesShouldBe)
}
