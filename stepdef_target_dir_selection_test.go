package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"strconv"
)

var targetDirSelected string
var isTargetValid bool

func userSelectsAsTarget(dir string) error {
	sourceDirSelected = dir
	return nil
}

func applicationStatusContains(message string) error {
	return godog.ErrPending
}


func validityOfTargetDirIs(validity string) error {
	var expectedValidity bool
	expectedValidity, err := strconv.ParseBool(validity)

	if err != nil {
		return fmt.Errorf( "error in converting %s to boolean value:", validity)
	}
	if expectedValidity != isTargetValid {
		return fmt.Errorf( "expected target directory to be %s but was %v", validity, isTargetValid)
	}
	return nil
}


func InitializeTargetDirSelectionScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		targetDirSelected = ""
		isTargetValid = false
	})

	ctx.Step(`^application status contains "([^"]*)"$`, applicationStatusContains)
	ctx.Step(`^user selects "([^"]*)" as target$`, userSelectsAsTarget)
	ctx.Step(`^validity of target dir is "([^"]*)"$`, validityOfTargetDirIs)

}
