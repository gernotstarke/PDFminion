package sample_pdfs

import (
	"github.com/stretchr/testify/assert"
	"pdfminion/domain"
	"testing"
)

// test if the example files exist
func TestSamplePDFs(t *testing.T) {

	// some valid PDF files shall be checked without error
	checkFor(t, OnePageFile)

	checkFor(t, ThreePageFile)

	// check if disguised markdown file exists
	valid, err := domain.ValidatePDFFile(DisguisedMarkdownFile)

	assert.NotNil(t, err)

	// disguised file shall NOT be valid!
	assert.Equal(t, false, valid)

}

func checkFor(t *testing.T, fileToCheck string) {
	_, err := domain.ValidatePDFFile(fileToCheck)

	if err != nil {
		t.Errorf("ValidatePDFFile: file %v: %v", fileToCheck, err)
	}

}
