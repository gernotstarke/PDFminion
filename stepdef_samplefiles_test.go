package main

import (
	"errors"
	"github.com/cucumber/godog"
	"pdfminion/fileutil"
)

var dirExists bool
var sampleDir string

func checkIfDirExists(dirName string) error {
	var err error

	sampleDir = dirName
	dirExists, err = fileutil.FileExists( dirName )

	if err != nil {
		return err
	}
	return nil
}

func theDirectoryContainingSampleFilesIsAvailable() error {

	if dirExists == true {
		return nil
	} else {
		return errors.New( "directory #{sampleDir} is not available")
	}
}

func InitializeSampleScenario(ctx *godog.ScenarioContext) {
	dirExists = false

	ctx.Step(`^Existence of "([^"]*)" is checked$`, checkIfDirExists)
	ctx.Step(`^The directory containing sample files is available$`, theDirectoryContainingSampleFilesIsAvailable)
}

