package main

import "github.com/cucumber/godog"

func anEmptyDirectory() error {
	return godog.ErrPending
}

func iHaveASourceDirectoryWithFileAndPages(arg1, arg2 int) error {
	return godog.ErrPending
}

func iStartNumbering() error {
	return godog.ErrPending
}

func isChecked(dirPrefix string, fname string) error {
	return godog.ErrPending
}

func isReturned(arg1 int) error {
	return godog.ErrPending
}

func numberOfFilesIsChecked() error {
	return godog.ErrPending
}

func shouldBe(arg1, arg2 int) error {
	return godog.ErrPending
}

func thePagecountShouldBe(arg1 int) error {
	return godog.ErrPending
}

func InitializeSamplePDFScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^an empty directory$`, anEmptyDirectory)
	ctx.Step(`^I have a sourceDirectory with (\d+) file and (\d+) pages$`, iHaveASourceDirectoryWithFileAndPages)
	ctx.Step(`^I start numbering$`, iStartNumbering)
	ctx.Step(`^(\d+) is checked$`, isChecked)
	ctx.Step(`^(\d+) (\d+) is checked$`, isChecked)
	ctx.Step(`^(\d+) is returned$`, isReturned)
	ctx.Step(`^Number of files is checked$`, numberOfFilesIsChecked)
	ctx.Step(`^(\d+) should be (\d+)$`, shouldBe)
	ctx.Step(`^the pagecount should be (\d+)$`, thePagecountShouldBe)
}


