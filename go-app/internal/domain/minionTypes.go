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
// Evenify
// Merge
// MergeFileName
// Debug
//
// For the page header:
// RunningHeader: The text to show in the page header, identical on all pages
//
// For the page footer:
// Let's take "Chapter 3 - page 3 of 42" as an example
// ChapterPrefix == "Chapter"
// Separator == " - "
// PagePrefix == "page"
// TotalPageCountPrefix == "of"

type MinionConfig struct {

	// Commands that are executed immediately
	Help          bool
	Version       bool
	ListLanguages bool

	// General config settings, see ADR-0005
	Language      language.Tag
	Debug         bool
	SourceDir     string
	TargetDir     string
	Force         bool
	Evenify       bool
	Merge         bool
	MergeFileName string

	// Page-related settings, see ADR-0006
	RunningHeader        string
	ChapterPrefix        string
	Separator            string
	PagePrefix           string
	TotalPageCountPrefix string
	BlankPageText        string
}

// DefaultTexts holds UI texts by language

var DefaultTexts = map[language.Tag]struct {
	ChapterPrefix string
	RunningHeader string
	PageFooter    string
	PageNumber    string
	BlankPageText string
}{
	language.German: {
		ChapterPrefix: "Kapitel",
		RunningHeader: "Seite",
		PageFooter:    "Seite %d von %d",
		PageNumber:    "Seite %d",
		BlankPageText: "Diese Seite bleibt absichtlich leer",
	},
	language.English: {
		ChapterPrefix: "Chapter",
		RunningHeader: "Page",
		PageFooter:    "Page %d of %d",
		PageNumber:    "Page %d",
		BlankPageText: "deliberately left blank",
	},
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
		Short:   "",
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
	"list-languages": {
		Long:    "list-languages",
		Short:   "ll",
		Default: false,
		Help:    "Show available languages, like EN or DE",
	},
	"defaults": {
		Long:    "defaults",
		Short:   "d",
		Default: false,
		Help:    "Show all current default settings",
	},
	"language": {
		Long:    "language",
		Short:   "l",
		Default: false,
		Help:    "Sets the language for stamped text, like EN or DE",
	}}

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
