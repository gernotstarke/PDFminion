package main

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"os"
	"pdfminion/domain"
	"pdfminion/sample_pdfs"
)

var pageCount int
var pdfFileName string

func aDirUnderSampleContainingFile(directory, samplePDFDir, pdfFile string) error {
	pdfFileName = samplePDFDir + string(os.PathSeparator) + directory + string(os.PathSeparator) + pdfFile

	fileExists, err := domain.FileExists(samplePDFDir)

	if (err != nil) || !fileExists {
		return fmt.Errorf("sample PDF file does not exist" + err.Error())
	} else {
		return nil
	}
}

func numberOfPagesIsCounted() error {
	pageCount, _ = domain.CountPagesOfPDFFile(pdfFileName)

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

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		pageCount = -1
		samplePDFDir = sample_pdfs.SampleDirectoryPrefix
		return ctx, nil
	})

	ctx.Step(`^A "([^"]*)" under "([^"]*)" containing "([^"]*)"$`, aDirUnderSampleContainingFile)
	ctx.Step(`^Number of pages is counted$`, numberOfPagesIsCounted)
	ctx.Step(`^the number of pages should be (\d+)$`, theNumberOfPagesShouldBe)
}
