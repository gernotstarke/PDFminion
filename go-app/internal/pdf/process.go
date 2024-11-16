package pdf

import (
	"fmt"
	"log"
	"pdfminion/go-app/internal/config"
	"sort"
)

func ProcessPDFs(cfg *config.Config) error {
	InitializePDFInternals()

	files, err := CollectCandidatePDFs(cfg)
	if err != nil {
		return fmt.Errorf("error collecting candidate PDFs: %w", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	pdfFiles, nrOfValidPDFs := ValidatePDFs(files)

	err = CopyValidatedPDFs(pdfFiles, cfg.SourceDir, cfg.TargetDir, cfg.Force)
	if err != nil {
		return fmt.Errorf("error during copy: %w", err)
	}

	log.Printf("%v", pdfFiles)

	Evenify(nrOfValidPDFs, pdfFiles)
	AddPageNumbersToAllFiles(nrOfValidPDFs, pdfFiles)

	return nil
}
