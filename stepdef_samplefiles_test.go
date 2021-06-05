package main

import (
	"errors"
	"github.com/cucumber/godog"
	"pdfminion/fileutil"
)

var dirExists bool
var testedDirectory string

func checkIfDirExists(dirName string) error {
	var err error

	testedDirectory = dirName
	dirExists, err = fileutil.FileExists( dirName )

	if err != nil {
		return err
	}
	return nil
}

// FixMe
func numberOfPDFFilesIsCounted() error {
	return godog.ErrPending
}

// FixMe
func isReturned(arg1 int) error {
	return godog.ErrPending
}

func setDirectory(dirName string) error {
	testedDirectory = dirName
	return nil
}


func theDirectoryContainingSampleFilesIsAvailable() error {

	if dirExists == true {
		return nil
	} else {
		return errors.New( "directory #{testedDirectory} is not available")
	}
}

func InitializeSampleScenario(ctx *godog.ScenarioContext) {
	dirExists = false

	ctx.Step(`^Existence of "([^"]*)" is checked$`, checkIfDirExists)
	ctx.Step(`^The directory containing sample files is available$`, theDirectoryContainingSampleFilesIsAvailable)

	ctx.Step(`^Directory "([^"]*)"$`, setDirectory)
	ctx.Step(`^(\d+) is returned$`, isReturned)
	ctx.Step(`^Number of PDF files is counted$`, numberOfPDFFilesIsCounted)

}

