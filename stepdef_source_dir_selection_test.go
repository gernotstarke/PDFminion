package main

import "github.com/cucumber/godog"

var sourceDirSelected string

func userSelectsAsSource(dir string) error {
	sourceDirSelected = dir
	return nil
}

func validityOfSourceDirIsTrue() error {
	return godog.ErrPending
}

func InitializeSourceDirSelectionScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		sourceDirSelected = ""
	})

	ctx.Step(`^user selects "([^"]*)" as source$`, userSelectsAsSource)
	ctx.Step(`^validity of source dir is true$`, validityOfSourceDirIsTrue)
}
