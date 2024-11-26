package domain

import (
	"golang.org/x/text/language"
)

type PDFConfig struct {
	Language      language.Tag
	SectionHeader string
	PageHeader    string
	PageFooter    string
	StartPage     int
	// ... other settings
}

// DefaultTexts holds UI texts by language

var DefaultTexts = map[language.Tag]struct {
	SectionHeader string
	PageHeader    string
	PageFooter    string
	PageNumber    string
}{
	language.German: {
		SectionHeader: "Kapitel",
		PageHeader:    "Seite",
		PageFooter:    "Seite %d von %d",
		PageNumber:    "Seite %d",
	},
	language.English: {
		SectionHeader: "Chapter",
		PageHeader:    "Page",
		PageFooter:    "Page %d of %d",
		PageNumber:    "Page %d",
	},
}
