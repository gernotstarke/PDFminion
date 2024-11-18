package pdf

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"io"
	"log"
	"os"
	"path/filepath"
	"pdfminion/internal/cli"
	"pdfminion/internal/util"
	"strconv"
)

type SingleFileToProcess struct {
	Filename      string
	PageCount     int
	OrigByteCount int64
}

var (
	// the relaxedConf is VERY specific to the pdfcpu library
	relaxedConf *model.Configuration
)

func InitializePDFInternals() {
	relaxedConf = model.NewDefaultConfiguration()
	relaxedConf.ValidationMode = model.ValidationRelaxed
}

func CollectCandidatePDFs(cfg *cli.Options) ([]string, error) {
	// count PDFs in source directory
	// abort if no PDF file is present

	var nrOfCandidatePDFs int

	files, err, nrOfCandidatePDFs := getNumberOfCandidatePDFs(cfg.SourceDir)

	// exit if no PDF files can be found
	if nrOfCandidatePDFs == 0 {
		fmt.Fprintf(os.Stderr, "error: no PDF files found in %s\n", cfg.SourceDir)
		os.Exit(1)
	}
	return files, err
}

func getNumberOfCandidatePDFs(sourceDir string) ([]string, error, int) {
	// collect all candidate PDFs with Glob
	// "candidate" means, PDF has not been validated
	pattern := filepath.Join(sourceDir, "*.pdf")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Println("Error:", err)
	}

	return files, err, len(files)
}

func ValidatePDFs(files []string) ([]SingleFileToProcess, int) {

	validPDFs := make([]SingleFileToProcess, 0)
	nrOfValidPDFs := 0

	for _, file := range files {
		err := api.ValidateFile(file, relaxedConf)
		if err != nil {
			log.Printf("%v is not a valid PDF, %v\n", file, err)
			continue
		}

		pageCount, err := api.PageCountFile(file)
		if err != nil {
			log.Printf("error counting pages in %v\n", file)
			continue
		}

		validPDFs = append(validPDFs, SingleFileToProcess{
			Filename:  filepath.Base(file),
			PageCount: pageCount,
		})
		nrOfValidPDFs++
	}

	return validPDFs, nrOfValidPDFs
}

func CopyValidatedPDFs(validPDFs []SingleFileToProcess, sourceDir, targetDir string, force bool) error {
	// Check if target directory is empty, unless force flag is set
	if !force {
		entries, err := os.ReadDir(targetDir)
		if err != nil {
			return fmt.Errorf("error reading target directory: %w", err)
		}
		if len(entries) > 0 {
			return fmt.Errorf("target directory is not empty")
		}
	}

	for i := range validPDFs {
		sourcePath := filepath.Join(sourceDir, validPDFs[i].Filename)
		targetPath := filepath.Join(targetDir, validPDFs[i].Filename)

		// Check if file exists and skip if not forcing overwrite
		if !force {
			if _, err := os.Stat(targetPath); err == nil {
				fmt.Printf("Skipping existing file: %s\n", targetPath)
				continue
			}
		}

		originalFile, err := os.Open(sourcePath)
		if err != nil {
			return fmt.Errorf("error opening source file %s: %w", sourcePath, err)
		}
		defer originalFile.Close()

		newFile, err := os.Create(targetPath)
		if err != nil {
			return fmt.Errorf("error creating target file %s: %w", targetPath, err)
		}
		defer newFile.Close()

		bytesWritten, err := io.Copy(newFile, originalFile)
		if err != nil {
			return fmt.Errorf("error copying file %s: %w", sourcePath, err)
		}

		validPDFs[i].OrigByteCount = bytesWritten
		fmt.Printf("Copied: %s\n", targetPath)

		// update the filename in the slice with the full path
		validPDFs[i].Filename = targetPath
	}

	return nil
}

func AddPageNumbersToAllFiles(nrOfValidPDFs int, pdfFiles []SingleFileToProcess) {
	// currentOffset is the _previous_ pagenumber
	var currentOffset = 0

	for i := 0; i < nrOfValidPDFs; i++ {
		var currentFilePageCount = pdfFiles[i].PageCount
		var currentFileName = pdfFiles[i].Filename
		log.Printf("File %s starts %d, ends %d", currentFileName, currentOffset+1,
			currentOffset+currentFilePageCount)

		err := api.AddWatermarksMapFile(currentFileName,
			"",
			watermarkConfigurationForFile(i+1,
				currentOffset,
				currentFilePageCount),
			relaxedConf)
		if err != nil {
			log.Println(err)
		}
		currentOffset += currentFilePageCount
	}
}

// create a map[int] of TextWatermark configurations
func watermarkConfigurationForFile(chapterNr, previousPageNr, pageCount int) map[int]*model.Watermark {

	wmcs := make(map[int]*model.Watermark)

	for page := 1; page <= (pageCount); page++ {
		var currentPageNr = previousPageNr + page
		var chapterStr = cli.ChapterPrefix + strconv.Itoa(chapterNr)
		var pageStr = cli.PageNrPrefix + strconv.Itoa(currentPageNr)

		wmcs[page], _ = api.TextWatermark(chapterStr+cli.ChapterPageSeparator+pageStr,
			waterMarkDescription(currentPageNr), true, false, types.POINTS)
	}
	return wmcs
}

const fontColorSize = "font:Helvetica, points:16, scale: 0.9 abs, rot: 0, color: 0.5 0.5 0.5"

// creates a pdfcpu TextWatermark description
func waterMarkDescription(pageNumber int) string {

	const evenPos string = "position: bl"
	const evenOffset string = "offset: 20 6"
	const oddPos string = "position: br"
	const oddOffset string = "offset: -20 6"

	positionAndOffset := ""

	if util.IsEven(pageNumber) {
		positionAndOffset = evenPos + "," + evenOffset
	} else {
		positionAndOffset = oddPos + "," + oddOffset
	}
	return fontColorSize + "," + positionAndOffset
}
