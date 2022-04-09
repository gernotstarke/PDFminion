package domain

import (
	samplePDFs "pdfminion/sample_pdfs"
	"testing"
)

const prefix = samplePDFs.SampleDirectoryPrefix

var testDirs = []struct {
	dirName          string
	expectedNrOfPDFs int
}{
	{ // empty Dir
		prefix + "EmptyDir",
		0,
	},
	{ // OnePDF
		prefix + "OnePDF",
		1,
	},
	{ // Four files, TWO PDFs
		prefix + "FourFilesTwoPDFs",
		2,
	},
	{ // 12 PDFfs
		prefix + "TwelvePDFs",
		12,
	},
	{ // One PDF, one subfolder with extension PDF that should not count
		prefix + "OnePDFWithSubfolder",
		1,
	},
}

func TestCountPDFsInDir(t *testing.T) {

	for _, d := range testDirs {
		got := CountPDFsInDir(d.dirName)
		if got != d.expectedNrOfPDFs {
			t.Errorf("FAIL: directory %s expected %v, got %v PDF files", d.dirName, d.expectedNrOfPDFs, got)
		}
	}

}
