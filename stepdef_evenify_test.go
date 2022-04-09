package main

import (
	"context"
	"github.com/cucumber/godog"
	"pdfminion/sample_pdfs"
)

func InitializeEvenifyScenario(ctx *godog.ScenarioContext) {

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		pageCount = -1
		samplePDFDir = sample_pdfs.SampleDirectoryPrefix
		return ctx, nil
	})

}
