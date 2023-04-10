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
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"pdfminion/domain"
	"sort"
	"strconv"
)

type singleFileToProcess struct {
	filename         string
	origPageCount    int
	hasBeenEvenified bool
}

const sourceDirName = "_pdfs"
const targetDirName = "_target"

// pdfFiles contains filenames, pagecounts, processing state etc.
//var pdfFiles []singleFileToProcess

func main() {

	domain.SetupConfiguration()

	// get current directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Printf("current directory: %s\n", currentDir)

	// count PDFs in current directory
	// abort, if no PDF file is present
	var nrOfCandidatePDFs int

	// collect all PDFs with Glob
	pattern := sourceDirName + "/*.pdf"
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Println("Error:", err)
	}

	nrOfCandidatePDFs = len(files)
	fmt.Printf("%d PDF files found.", nrOfCandidatePDFs)

	// exit if no PDF files can be found
	if nrOfCandidatePDFs == 0 {
		fmt.Fprintf(os.Stderr, "error: no PDF files found\n")
		os.Exit(1)
	}

	// sort files alphabetically (as we cannot assume any sort order from `os.Glob)
	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	// create target directory
	// TODO: check if files already present in target directory
	if _, err := os.Stat(targetDirName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(targetDirName, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// create slice of singleFileToProcess of required length
	pdfFiles := make([]singleFileToProcess, nrOfCandidatePDFs)

	// initialize slice of singleFileToProcess
	// move over the pdf files into pdfFiles variable

	// TODO: insert only valid PDFs in pdfFiles, count accordingly

	var originalFile, newFile *os.File

	var nrOfValidPDFs = 0
	for i := 0; i < nrOfCandidatePDFs; i++ {

		// check if file-i is a valid PDF with pdfcpu.api
		// use default configuration for pdfcpu ("nil")
		err = api.ValidateFile(files[i], nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v is no valid PDF\n", files[i])
		} else {
			// we have a valid PDF

			nrOfValidPDFs++

			// count the pages of this particular file
			// TODO: handle zero-page PDFs
			pdfFiles[i].origPageCount, err = api.PageCountFile(files[i])

			if err != nil {
				fmt.Fprintf(os.Stderr, "error counting pages in %v\n", files[i])
			} else {

				// create target filePath
				pdfFiles[i].filename = filepath.Join(targetDirName, filepath.Base(files[i]))

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

				//This will copy
				bytesWritten, err := io.Copy(newFile, originalFile)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Bytes Written: %d\n", bytesWritten)

			}

		}
	}

	fmt.Printf("%v", pdfFiles)

	// TODO: add page numbers

	// currentOffset is the _previous_ pagenumber
	var currentOffset = 0

	for i := 0; i < nrOfCandidatePDFs; i++ {
		var currentFilePageCount = pdfFiles[i].origPageCount
		log.Printf("File %s starts %d, ends %d", pdfFiles[i].filename, currentOffset,
			currentOffset+currentFilePageCount)

		err := api.AddWatermarksMapFile(pdfFiles[i].filename,
			"",
			watermarkConfigurationForFile(currentOffset,
				currentFilePageCount),
			nil)
		if err != nil {
			log.Println(err)
		}
		currentOffset += currentFilePageCount
	}

}

// create a map[int] of TextWatermark configurations
func watermarkConfigurationForFile(previousPageNr, pageCount int) map[int]*pdfcpu.Watermark {

	wmcs := make(map[int]*pdfcpu.Watermark)

	for page := 1; page <= (pageCount); page++ {
		var currentPageNr = previousPageNr + page
		wmcs[page], _ = api.TextWatermark(strconv.Itoa(currentPageNr),
			waterMarkDescription(page), true, false, pdfcpu.POINTS)
	}
	return wmcs
}

// creates a pdfcpu TextWatermark description
func waterMarkDescription(pageNumber int) string {
	fontColorSize := "font:Helvetica, points:12, rot: 0, scale:0.05, color: 0.5 0.5 0.5"

	const evenPos string = "position: bl"
	const evenOffset string = "offset: 0 0"
	const oddPos string = "position: br"
	const oddOffset string = "offset: 0 0"

	positionAndOffset := ""

	if isEven(pageNumber) {
		positionAndOffset = evenPos + "," + evenOffset
	} else {
		// fmt.Println(pageNumber,"is odd")
		positionAndOffset = oddPos + "," + oddOffset
	}
	return fontColorSize + "," + positionAndOffset
}

func isEven(nr int) bool {
	if nr%2 == 0 {
		return true
	} else {
		return false
	}
}
