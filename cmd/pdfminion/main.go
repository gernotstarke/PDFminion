package main

import (
	"fmt"
	"os"
	"pdfminion/internal/config"
	"pdfminion/internal/pdf"
)

func main() {
	cfg := config.New()
	cfg.ParseFlags()

	if err := cfg.Evaluate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := pdf.ProcessPDFs(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing PDFs: %v\n", err)
		os.Exit(1)
	}
}
