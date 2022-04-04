package main

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"os"
	"pdfminion/fileutil"
	"pdfminion/sample_pdfs"
	"strings"
)

var (
	pdfFilesInDir int
	samplePDFDir  string
	pdfFileList   []string
)

func aDirectoryNamedContainingSamplePDFs(sampleDir string) error {
	// TODO: fix redundant definition of samplePDFDir
	samplePDFDir = sampleDir
	return nil
}

func numberOfPDFFilesInIsCounted(dir string) error {
	var thisDir = samplePDFDir + string(os.PathSeparator) + dir
	if thisDir != "" {
		nrOfPDFFiles = fileutil.CountPDFsInDir(thisDir)
	} else {
		return fmt.Errorf("no directory given (dir == #{thisDir}")
	}
	return nil
}

// ensures sample PDF directory exists
func samplePDFFilesUnder(dir string) error {

	dirExists, err := fileutil.FileExists(samplePDFDir)

	if (err != nil) || !dirExists {
		return fmt.Errorf("standard sample PDF dir does not exist" + err.Error())
	} else {
		return nil
	}
}

func theNumberOfPDFFilesShouldBe(expectedNrOfPDFFiles int) error {
	if expectedNrOfPDFFiles != nrOfPDFFiles {
		return fmt.Errorf("expected %d PDF files, but found %d", expectedNrOfPDFFiles, nrOfPDFFiles)
	} else {
		return nil
	}
}

func pDFFilesInDirAreCollected(dir string) error {
	var thisDir = samplePDFDir + string(os.PathSeparator) + dir
	if thisDir != "" {
		nrOfPDFFiles, pdfFileList = fileutil.CountAndCollectPDFsInDir(thisDir)
	} else {
		return fmt.Errorf("no directory given (dir == #{thisDir}")
	}
	return nil
}

func theListOfPDFFilesShallBe(expectedFiles string) error {
	if expectedFiles != strings.Join(pdfFileList, " ") {
		return fmt.Errorf("expected %s PDF files, but found %s", expectedFiles, pdfFileList)
	} else {
		return nil
	}

	return nil
}

func InitializeFileHandlingScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		pageCount = -1
		samplePDFDir = sample_pdfs.SampleDirectoryPrefix
		return ctx, nil
	})

	ctx.Step(`^A directory named "([^"]*)" containing sample PDFs$`, aDirectoryNamedContainingSamplePDFs)

	ctx.Step(`^Sample PDF files in "([^"]*)"$`, samplePDFFilesUnder)
	ctx.Step(`^PDF files in "([^"]*)" are counted$`, numberOfPDFFilesInIsCounted)

	ctx.Step(`^PDF files in "([^"]*)" are collected$`, pDFFilesInDirAreCollected)
	ctx.Step(`^the list of PDF files shall be "([^"]*)"$`, theListOfPDFFilesShallBe)

	ctx.Step(`^the number of PDF files should be (\d+)$`, theNumberOfPDFFilesShouldBe)

}
