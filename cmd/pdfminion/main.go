package main

/**
minimal viable product version of PDFminion:
* works in the currently active directory
* no config parameters
* numbers all PDFs present in directory

*/
import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"io"
	"log"
	"os"
	"path/filepath"
	"pdfminion/internal/config"
	"pdfminion/internal/util"
	"sort"
	"strconv"
)

type singleFileToProcess struct {
	filename      string
	pageCount     int
	origByteCount int64
}

const blankPageNote = "Diese Seite bleibt\n absichtlich frei"
const pageNrPrefix = ""
const chapterPrefix = "Kap."
const chapterPageSeparator = " - "

func main() {
	cfg := config.New()
	cfg.ParseFlags()

	if err := cfg.Evaluate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// count PDFs in current directory
	// abort, if no PDF file is present
	var nrOfCandidatePDFs int

	// collect all candidate PDFs with Glob
	// "candidate" means, PDF has not been validated
	pattern := filepath.Join(cfg.SourceDir, "*.pdf")
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Println("Error:", err)
	}

	nrOfCandidatePDFs = len(files)

	// exit if no PDF files can be found
	if nrOfCandidatePDFs == 0 {
		fmt.Fprintf(os.Stderr, "error: no PDF files found in %s\n", cfg.SourceDir)
		os.Exit(1)
	}

	// sort files alphabetically (as we cannot assume any sort order from `os.Glob)
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// create slice of singleFileToProcess of required length
	pdfFiles := make([]singleFileToProcess, nrOfCandidatePDFs)

	// initialize slice of singleFileToProcess
	// move over only the validated pdf files into pdfFiles variable

	var originalFile, newFile *os.File

	// relax the validation mode, otherwise asciidoc generated pdfs
	// will break and lead to errors
	// see https://github.com/pdfcpu/pdfcpu/issues/498

	relaxedConf := model.NewDefaultConfiguration()
	relaxedConf.ValidationMode = model.ValidationRelaxed

	var nrOfValidPDFs = 0
	for i := 0; i < nrOfCandidatePDFs; i++ {

		// check if file-i is a valid PDF with pdfcpu.api
		// use default configuration for pdfcpu ("nil")
		err = api.ValidateFile(files[i], relaxedConf)
		if err != nil {
			log.Printf("%v is no valid PDF, %v\n", files[i], err)
		} else {

			// we have a valid PDF

			nrOfValidPDFs++

			// count the pages of this particular file
			// TODO: handle zero-page PDFs
			pdfFiles[i].pageCount, err = api.PageCountFile(files[i])

			if err != nil {
				log.Printf("error counting pages in %v\n", files[i])
			} else {

				// create target filePath
				pdfFiles[i].filename = filepath.Join(cfg.TargetDir, filepath.Base(files[i]))

				// copy that particular file to _target
				// Open original file
				originalFile, err = os.Open(files[i])
				if err != nil {
					log.Fatal(err)
				}
				defer originalFile.Close()

				// Create new file
				newFile, err = os.Create(pdfFiles[i].filename)
				if err != nil {
					log.Fatal(err)
				}
				defer newFile.Close()

				//This will copy.
				bytesWritten, err := io.Copy(newFile, originalFile)
				if err != nil {
					log.Fatal(err)
				}
				pdfFiles[i].origByteCount = bytesWritten
			}

		}
	}

	log.Printf("%v", pdfFiles)

	// evenify: add empty page to every file with even pagecount
	for i := 0; i < nrOfValidPDFs; i++ {
		if !util.IsEven(pdfFiles[i].pageCount) {
			// add single blank page at the end of the file
			_ = api.InsertPagesFile(pdfFiles[i].filename, "", []string{strconv.Itoa(pdfFiles[i].pageCount)}, false, relaxedConf)

			pdfFiles[i].pageCount++

			// TODO: add huge diagonal marker text "deliberately left blank" to new blank page

			onTop := true
			update := false

			wm, err := api.TextWatermark(blankPageNote, "font:Helvetica, points:48, col: 0.5 0.6 0.5, rot:45, sc:1 abs",
				onTop, update, types.POINTS)
			if err != nil {
				log.Println("Error creating watermark configuration %v: %v", wm, err)
			} else {

				err = api.AddWatermarksFile(pdfFiles[i].filename, "", []string{strconv.Itoa(pdfFiles[i].pageCount)}, wm,
					relaxedConf)

				if err != nil {
					log.Println("error stamping blank page in file %v: %v", pdfFiles[i].filename, err)
				}

			}
			log.Println("File %s was evenified", pdfFiles[i].filename)
		}
	}

	// add page numbers

	// currentOffset is the _previous_ pagenumber
	var currentOffset = 0

	for i := 0; i < nrOfValidPDFs; i++ {
		var currentFilePageCount = pdfFiles[i].pageCount
		var currentFileName = pdfFiles[i].filename
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
		var chapterStr = chapterPrefix + strconv.Itoa(chapterNr)
		var pageStr = pageNrPrefix + strconv.Itoa(currentPageNr)

		wmcs[page], _ = api.TextWatermark(chapterStr+chapterPageSeparator+pageStr,
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
