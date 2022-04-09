package domain

import (
	"errors"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CheckSrcDirectorStatus verifies the directory exists and contains PDF files.
// It sets the status-info accordingly.
// In case of errors, an appropriate status is set in the app configuration.
func CheckSrcDirectoryStatus(dirName string) bool {
	log.Printf("ChedkSrcDirStatus called")
	if !IsDirExists(dirName, "source") {
		// file does not exist
		SetStatusInfo(dirName + " does not exist.")
		SetSourceDirMessage("Select source directory.")
		return false
	} else { // err == nil
		// directory exists, now check for PDF files
		nrOfPDFs := CountPDFsInDir(dirName)

		if nrOfPDFs == 0 {
			SetStatusInfo(dirName + " has no PDF files.")
			SetSourceDirMessage("0 (zero) PDF files found.")
			return false
		} else { // nrOfPDFs > 0
			SetSourceDirMessage(strconv.Itoa(nrOfPDFs) + "PDF files found")
			SetStatusInfo("source directory ok")
			return true
		}
	}
}

// IsDirExists returns true if the directory exists, false otherwise
func IsDirExists(dirName string, dirType string) bool {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return false
	} else {
		SetStatusInfo(dirName + ": other error")
		return false
	}
	return true
}

// FileExists checks if a file exists.
// It does not care for file extensions, type, path or size.
func FileExists(fileName string) (bool, error) {
	var err error

	// Stat returns file info. It will return
	// an error if there is no file.
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

func GetUserHomeDirectory() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
