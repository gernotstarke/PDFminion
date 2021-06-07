package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"pdfminion/fileutil"
)

var pdfFilesInDir int
var samplePDFDir string

func numberOfPDFFilesInIsCounted(dir string) error {
	var thisDir string = samplePDFDir + "/" + dir
	if thisDir != "" {
		nrOfPDFFiles = fileutil.CountPDFsInDir(thisDir)
	} else {
		return fmt.Errorf( "no directory given (dir == #{thisDir}")
	}
	return nil
}

// ensures sample PDF directory exists
func samplePDFFilesUnder(dir string) error {

	samplePDFDir = dir

	dirExists, err := fileutil.FileExists(samplePDFDir)

	if (err != nil) || !dirExists {
		return fmt.Errorf("standard sample PDF dir does not exist" + err.Error())
	} 	else  {
		return nil
	}
}

func theNumberOfPDFFilesShouldBe(expectedNrOfPDFFiles int) error {
	if expectedNrOfPDFFiles != nrOfPDFFiles {
		return fmt.Errorf("expected %d PDF files, but found %d", expectedNrOfPDFFiles, nrOfPDFFiles)
	} else {
		return nil
	}
}

func aFileCanBeCreatedThere() error {
	return godog.ErrPending
}

func aTemporaryDirectoryIsCreated() error {
	return godog.ErrPending
}

func InitializeFileHandlingScenario(ctx *godog.ScenarioContext) {


	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		pdfFilesInDir = -1
	})

	ctx.Step(`^Sample PDF files under "([^"]*)"$`, samplePDFFilesUnder)
	ctx.Step(`^Number of PDF files in "([^"]*)" is counted$`, numberOfPDFFilesInIsCounted)
	ctx.Step(`^the number of PDF files should be (\d+)$`, theNumberOfPDFFilesShouldBe)

	ctx.Step(`^a temporary directory is created$`, aTemporaryDirectoryIsCreated)
	ctx.Step(`^a file can be created there$`, aFileCanBeCreatedThere)

}
