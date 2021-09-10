package fileutil

import (
	"errors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// FileExists checks if a file exists.
// It does not care for file extensions, type, path or size.
func FileExists(fileName string) (bool, error) {
	var err error

	// os.Stat returns file info.
	// It will return an error if there is no file.
	_, err = os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, err
		}
	}
	return true, nil
}

// extensionValid checks if file has required extension.
func extensionValid(fileName string, needExtension string) bool {
	extension := strings.TrimLeft(strings.ToUpper(filepath.Ext(fileName)), ".")
	return extension == needExtension
}

// ValidatePDFFile checks if file is OK: does it exist, does it have the proper extension, is it PDF?
func ValidatePDFFile(pdfFileName string) (bool, error) {

	// check if file exists
	_, err := FileExists(pdfFileName)
	if err != nil {
		log.Printf("File %v does not exist.", pdfFileName)
		return false, err
	}

	// check if file has proper extension PDF
	if !extensionValid(pdfFileName, "PDF") {
		msg := "File" + pdfFileName + " has wrong extension"
		log.Println(msg)
		return false, errors.New(msg)
	}

	// now validate PDF itself
	// use default configuration for pdfcpu ("nil") to validate PDF itself
	err = api.ValidateFile(pdfFileName, nil)
	if err != nil {
		msg := "File" + pdfFileName + " is no valid PDF."
		log.Println(msg)
		return false, errors.New(msg)
	}

	return true, nil
}


// 