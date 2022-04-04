package main

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"pdfminion/fileutil"
)

var dirExists bool
var testedDirectory string
var nrOfPDFFiles int

func checkIfDirExists(dirName string) error {
	var err error

	testedDirectory = dirName
	dirExists, err = fileutil.FileExists(dirName)

	if err != nil {
		return err
	}
	return nil
}

func numberOfPDFFilesIsCounted() error {

	if testedDirectory != "" {
		nrOfPDFFiles = fileutil.CountPDFsInDir(testedDirectory)
	} else {
		return fmt.Errorf("no directory given (testedDirectory == #{testedDirectory}")
	}
	return nil
}

//
func isReturned(expectedNrOfFiles int) error {

	if nrOfPDFFiles == expectedNrOfFiles {
		return nil
	} else {
		return fmt.Errorf(
			"expected nr of PDF files %d but it is %d",
			expectedNrOfFiles,
			nrOfPDFFiles,
		)

	}

}

func setDirectory(dirName string) error {
	testedDirectory = dirName
	return nil
}

func theDirectoryContainingSampleFilesIsAvailable() error {

	if dirExists == true {
		return nil
	} else {
		return errors.New("directory #{testedDirectory} is not available")
	}
}

func InitializeSampleScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		dirExists = false
		testedDirectory = ""
		nrOfPDFFiles = -1
	})

	ctx.Step(`^Existence of "([^"]*)" is checked$`, checkIfDirExists)
	ctx.Step(`^The directory containing sample files is available$`, theDirectoryContainingSampleFilesIsAvailable)

	ctx.Step(`^Directory "([^"]*)"$`, setDirectory)
	ctx.Step(`^(\d+) is returned$`, isReturned)
	ctx.Step(`^Number of PDF files is counted$`, numberOfPDFFilesIsCounted)

}
