package pdf

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"pdfminion/internal/domain"
	"sort"
)

func ProcessPDFs(cfg *domain.MinionConfig) error {
	log.Debug().Msg("Starting PDF processing") // Only shown in debug mode

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

	log.Debug().Int("fileCount", len(files)).Msg("Found files")

	Evenify(nrOfValidPDFs, pdfFiles)
	AddPageNumbersToAllFiles(nrOfValidPDFs, pdfFiles)

	return nil
}
