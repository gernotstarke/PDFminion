package main

/**
minimal viable product version of PDFminion:
* works in the currently active directory
* no config parameters
* numbers all PDFs present in directory

*/
import (
	"fmt"
	"log"
	"os"
	"pdfminion/internal/config"
	"pdfminion/internal/pdf"
	"sort"
)

func main() {

	// TODO: simplify cfg handling -> put all flag handling (parsing, eval + error handling) in config package
	cfg := config.New()
	cfg.ParseFlags()

	if err := cfg.Evaluate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// some technical stuff needs to be initialized before we can handle PDF files
	pdf.InitializePDFInternals()

	files, err := pdf.CollectCandidatePDFs(cfg)

	// sort, validate, copy

	// sort files alphabetically (as we cannot assume any sort order from `os.Glob)
	// TODO: move sort to processFiles!!
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// ensure we process only valid PDFs
	// * correct file format
	// * at least one page
	// TODO: handle zero-page PDFs
	// TODO: validatePDFs should return error if error...
	pdfFiles, nrOfValidPDFs := pdf.ValidatePDFs(files)

	// copy all valid PDFs to target directory
	// update pdfFiles slice to contain full path to target files
	err = pdf.CopyValidatedPDFs(pdfFiles, cfg.SourceDir, cfg.TargetDir, cfg.Force)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during copy, aborting %v\n", err)
		os.Exit(1)
	}

	// TODO: log only in DEBUG mode
	log.Printf("%v", pdfFiles)

	// Evenify: add empty page to every file with even page count
	pdf.Evenify(nrOfValidPDFs, pdfFiles)

	// add page numbers
	pdf.AddPageNumbersToAllFiles(nrOfValidPDFs, pdfFiles)

}
