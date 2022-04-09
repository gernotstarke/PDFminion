package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"pdfminion/domain"
)

func fileIsPresent(fname string) error {

	makeExist, err := domain.FileExists(fname)

	if (err != nil) || (makeExist == false) {
		return fmt.Errorf("error in looking for %s", fname)
	}
	return nil
}

func theRepositoryIsCheckedOut() error {
	return nil
}

func weHaveAFile(arg1 string) error {
	return godog.ErrPending
}

func InitializeMakefileScenario(ctx *godog.ScenarioContext) {

	ctx.Step(`^"([^"]*)" is present$`, fileIsPresent)
	ctx.Step(`^The repository is checked out$`, theRepositoryIsCheckedOut)
	ctx.Step(`^we have a "([^"]*)" file$`, weHaveAFile)

}
