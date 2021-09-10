package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
)

var sourceDirSelected string
var isSourceValid bool

func userSelectsAsSource(dir string) error {
	sourceDirSelected = dir
	return nil
}

func validityOfSourceDirIs(validity string) error {
	var expectedValidity bool
	expectedValidity, err := strconv.ParseBool(validity)

	if err != nil {
		return fmt.Errorf( "error in converting %s to boolean value:", validity)
	}

	// check if sourceDir is valid

	if expectedValidity != isSourceValid {
		return fmt.Errorf( "expected source directory to be %s but was %v", validity, isSourceValid)
	}
	return nil
}


func InitializeSourceDirSelectionScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		sourceDirSelected = ""
		isSourceValid = false
	})

	ctx.Step(`^user selects "([^"]*)" as source$`, userSelectsAsSource)
	ctx.Step(`^validity of source dir is "([^"]*)"$`, validityOfSourceDirIs)

}
