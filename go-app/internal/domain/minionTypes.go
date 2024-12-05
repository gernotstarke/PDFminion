package domain

import (
	"golang.org/x/text/language"
)

// MinionConfig holds configuration items
// -------------------------------------
// General settings:
// Language: The language to use for the text written onto the PDF
// SourceDir
// TargetDir
// Force
// Debug
//
// For the page header:
// RunningHead: The text to show in the page header, identical on all pages
//
// For the page footer:
// Let's take "Chapter 3 - page 3 of 42" as an example
// ChapterPrefix == "Chapter"
// Separator == " - "
// PagePrefix == "page"
// TotalPageCountPrefix == "of"

type MinionConfig struct {

	// General config settings, see ADR-0005
	Language  language.Tag
	Debug     bool
	SourceDir string
	TargetDir string
	Force     bool

	// Page-related settings, see ADR-0006
	RunningHead          string
	ChapterPrefix        string
	Separator            string
	PagePrefix           string
	TotalPageCountPrefix string
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

type Options struct {
	SourceDir string
	TargetDir string
	Force     bool
	Debug     bool
	Help      bool
	Version   bool
}

// FlagDef represents a flag definition with possible short and long forms
type FlagDef struct {
	Long    string
	Short   string
	Default interface{}
	Help    string
}

// Flag definitions
var Flags = map[string]FlagDef{
	"source": {
		Long:    "source",
		Short:   "s",
		Default: DefaultSourceDir,
		Help:    "Specify the source directory",
	},
	"target": {
		Long:    "target",
		Short:   "t",
		Default: DefaultTargetDir,
		Help:    "Specify the target directory",
	},
	"force": {
		Long:    "force",
		Short:   "f",
		Default: false,
		Help:    "Forces overwrite of existing files",
	},
	"debug": {
		Long:    "debug",
		Short:   "d",
		Default: false,
		Help:    "Enable debug logging",
	},
	"help": {
		Long:    "help",
		Short:   "h",
		Default: false,
		Help:    "Show this help message",
	},
	"version": {
		Long:    "version",
		Short:   "v",
		Default: false,
		Help:    "Show version information",
	},
}

const (
	DefaultSourceDir     = "_pdfs"
	DefaultTargetDir     = "_target"
	PageNrPrefix         = ""
	ChapterPrefix        = "Kap."
	ChapterPageSeparator = " - "
)

var AppVersion string

func SetAppVersion(version string) {
	AppVersion = version
}
