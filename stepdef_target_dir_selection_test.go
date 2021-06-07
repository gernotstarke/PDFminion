package main

import "github.com/cucumber/godog"

var targetDirSelected string

func userSelectsAsTarget(dir string) error {
	sourceDirSelected = dir
	return nil
}


func InitializeTargetDirSelectionScenario(ctx *godog.ScenarioContext) {

	// runs before a scenario is tested
	ctx.BeforeScenario(func(*godog.Scenario) {
		targetDirSelected = ""
	})

	ctx.Step(`^user selects "([^"]*)" as target$`, userSelectsAsTarget)
}
