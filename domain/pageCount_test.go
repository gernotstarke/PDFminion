package domain

import (
	samplePDFs "pdfminion/sample_pdfs"
)

var testPDFs = []struct {
	fileName          string
	expectedNrOfPages int
	expectedError     bool
}{
	{ // ONE page PDF
		samplePDFs.OnePageFile,
		1,
		false,
	},
	{ // THREE pages PDF
		samplePDFs.ThreePageFile,
		3,
		false,
	},
}

/*
func TestPageCount(t *testing.T) {
	want := 1
	fmt.Println(samplePDFs.OnePageFile)

	if got, _ := CountPagesOfPDFFile(samplePDFs.OnePageFile); got != want {
		t.Errorf("FAIL: CountPagesOfPDFFile() = %v, want %v", got, want)
	}

	if got, _ := CountPagesOfPDFFile(samplePDFs.ThreePageFile); got != 3 {
		t.Errorf("FAIL: CountPagesOfPDFFile()  %v, expected %v pages, got %v", samplePDFs.ThreePageFile, 3, got)
	}

	for _, f := range testPDFs {
		got, err := CountPagesOfPDFFile(f.fileName)

		if (err == nil) == f.expectedError {
			t.Errorf("FAIL: file %s should yield error, but didnt", f.fileName)
		}

		if got != f.expectedNrOfPages {
			t.Errorf("FAIL: file %s expected %v pages, got %v", f.fileName, f.expectedNrOfPages, got)
		}
	}
*/
