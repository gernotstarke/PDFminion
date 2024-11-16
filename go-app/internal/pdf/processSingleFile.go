package pdf

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"log"
	"pdfminion/go-app/internal/util"
	"strconv"
)

const blankPageNote = "Diese Seite bleibt\n absichtlich frei"

// TODO: Split Evenify function into loop-over-all-files and evenify-single-file

func Evenify(nrOfValidPDFs int, pdfFiles []SingleFileToProcess) {
	relaxedConf := model.NewDefaultConfiguration()
	relaxedConf.ValidationMode = model.ValidationRelaxed

	for i := 0; i < nrOfValidPDFs; i++ {
		if !util.IsEven(pdfFiles[i].PageCount) {
			// add single blank page at the end of the file
			_ = api.InsertPagesFile(pdfFiles[i].Filename, "", []string{strconv.Itoa(pdfFiles[i].PageCount)}, false, relaxedConf)

			pdfFiles[i].PageCount++

			onTop := true
			update := false

			wm, err := api.TextWatermark(blankPageNote, "font:Helvetica, points:48, col: 0.5 0.6 0.5, rot:45, sc:1 abs",
				onTop, update, types.POINTS)
			if err != nil {
				log.Printf("Error creating watermark configuration %v: %v\n", wm, err)
			} else {

				err = api.AddWatermarksFile(pdfFiles[i].Filename, "", []string{strconv.Itoa(pdfFiles[i].PageCount)}, wm,
					relaxedConf)

				if err != nil {
					log.Printf("error stamping blank page in file %v: %v\n", pdfFiles[i].Filename, err)
				}

			}
			log.Printf("File %s was evenified\n", pdfFiles[i].Filename)
		}
	}
}
