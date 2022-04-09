package domain

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CountPDFsInDir returns the number of PDF files in the given directory
func CountPDFsInDir(dirName string) int {

	nrOfPDFFiles, _ := CountAndCollectPDFsInDir(dirName)
	return nrOfPDFFiles
}

// CountAndCollectPDFsInDir returns the number of PDF files in the given directory
func CountAndCollectPDFsInDir(dirName string) (int, []string) {

	var files []string
	var fCount int

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {

		// this error checking is required to avoid panic
		if err != nil {
			return err
		}

		// exclude dirs
		if info.IsDir() {
			return nil
		}

		// count and collect only PDF files
		if strings.ToUpper(filepath.Ext(path)) == ".PDF" {
			fCount++
			files = append(files, info.Name())
		}
		return nil
	})

	if err != nil {
		log.Println("Error in walking the filepath " + dirName)
	}

	return fCount, files
}
