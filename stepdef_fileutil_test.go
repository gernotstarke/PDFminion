package main

import "github.com/cucumber/godog"

func aDifferentTargetDirectory() error {
	return godog.ErrPending
}

func aDirectoryContainingASingleFile() error {
	return godog.ErrPending
}

func aSourceDirectory() error {
	return godog.ErrPending
}

func anEmptyDirectoryIsSelectedAsSourcedir() error {
	return godog.ErrPending
}

func numberingIsDoneInAlphabeticalOrderOfTheFilenames() error {
	return godog.ErrPending
}

func pDFFilesAreCounted() error {
	return godog.ErrPending
}

func sequentialPageNumbersShouldHaveBeenAddedToThePDFFiles() error {
	return godog.ErrPending
}

func sequentialPageNumbersShouldHaveBeenAddedToThisPDFFileStartingFrom(arg1 int) error {
	return godog.ErrPending
}

func theNumberOfPDFFilesShouldBe(arg1 int) error {
	return godog.ErrPending
}

func theSourceDirectoryContainsASinglePDFFile() error {
	return godog.ErrPending
}

func theSourceDirectoryContainsSeveralPDFFiles() error {
	return godog.ErrPending
}

func thisFileHasTypePDF() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a different target directory$`, aDifferentTargetDirectory)
	ctx.Step(`^A directory containing a single file$`, aDirectoryContainingASingleFile)
	ctx.Step(`^A source directory$`, aSourceDirectory)
	ctx.Step(`^An empty directory is selected as source-dir$`, anEmptyDirectoryIsSelectedAsSourcedir)
	ctx.Step(`^numbering is done in alphabetical order of the filenames$`, numberingIsDoneInAlphabeticalOrderOfTheFilenames)
	ctx.Step(`^PDF files are counted$`, pDFFilesAreCounted)
	ctx.Step(`^sequential page numbers should have been added to the PDF files$`, sequentialPageNumbersShouldHaveBeenAddedToThePDFFiles)
	ctx.Step(`^sequential page numbers should have been added to this PDF file starting from (\d+)$`, sequentialPageNumbersShouldHaveBeenAddedToThisPDFFileStartingFrom)
	ctx.Step(`^the number of PDF files should be (\d+)$`, theNumberOfPDFFilesShouldBe)
	ctx.Step(`^the source directory contains a single PDF file$`, theSourceDirectoryContainsASinglePDFFile)
	ctx.Step(`^the source directory contains several PDF files$`, theSourceDirectoryContainsSeveralPDFFiles)
	ctx.Step(`^this file has type PDF$`, thisFileHasTypePDF)
}
