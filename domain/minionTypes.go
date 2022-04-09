package domain

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"log"
	"os"
)

var config MinionConfiguration

type MinionConfiguration struct {

	// MinionConfiguration represents the "state" of the application -
	// which files/directories are  selected and which processing options are configured,
	// plus status messages that might be displayed in some UI

	// directories
	// ************************************************
	sourceDirectory os.File
	targetDirectory os.File

	sourceDirName string
	targetDirName string

	// footer configuration settings
	// ************************************************

	numerate             bool
	pageNumberPrefix     string
	chapterPageSeparator string
	chapterPrefix        string

	// status information (in UI: shown in status line at bottom of UI)
	statusInfo       string
	sourceDirMessage string
	targetDirMessage string

	// header configuration
	// ************************************************
	headerText string

	// additional processing options
	// ************************************************

	// evenify: make every PDF have an EVEN page count,
	// by adding a single blank page to those PDFs with
	// originally odd number of pages
	evenify bool

	// this String gets stamped onto pages added during evenification
	blankPageText string "Diese Seite bleibt absichtlich frei"

	// concatenate several PDFs to a single PDF
	concatenate bool

	// if we concatenate and/or pad, how shall the resulting file be called
	resultingPDFFileName string
} // MinionConfiguration

// getter functions for better data encapsulation
// ==============================================
func StatusInfo() string           { return config.statusInfo }
func SourceDirName() string        { return config.sourceDirName }
func SourceDirMessage() string     { return config.sourceDirMessage }
func TargetDirName() string        { return config.targetDirName }
func TargetDirMessage() string     { return config.targetDirMessage }
func PageNumberPrefix() string     { return config.pageNumberPrefix }
func ChapterPageSeparator() string { return config.chapterPageSeparator }
func ChapterPrefix() string        { return config.chapterPrefix }
func HeaderText() string           { return config.headerText }
func IsEvenify() bool              { return config.evenify }
func BlankPageText() string        { return config.blankPageText }
func IsConcatenate() bool          { return config.concatenate }

// setter functions to avoid uncontrolled changes to global data...
// ================
func SetStatusInfo(msg string) {
	config.statusInfo = msg
}

func SetSourceDirMessage(msg string) {
	config.sourceDirMessage = msg
}

// SetSourceDirName also sets status messages
func SetSourceDirName(srcDir string) {
	log.Println("SetSourceDirName called")
	if CheckSrcDirectoryStatus(srcDir) {
		config.sourceDirName = srcDir
	} else {
		SetSourceDirMessage("illegal source dir")
	}
}

func SetTargetDirName(targetDir string) error                   { return nil }
func SetNumerate(numerate bool) error                           { return nil }
func SetPageNumberPrefix(pageNumberPrefix string) error         { return nil }
func SetChapterPageSeparator(chapterPageSeparator string) error { return nil }
func SetChapterPrefix(chapterPrefix string) error               { return nil }
func SetHeaderText(headerText string) error                     { return nil }
func SetEvenify(evenify bool) error                             { return nil }
func SetBlankPageText(blankPageText string) error               { return nil }
func SetConcatenate(concatenate bool) error                     { return nil }

// todo: use constructor function instead
func SetupConfiguration() {

	setupLanguageNeutralConfig()

	switch lang := checkPreferredLanguage(); lang {
	case "German":
		log.Println("Deutsch als Sprache identifiziert.")
		setupDEConfig()
	case "English":
		// fmt.Println("English identified as user language.")
		setupENConfig()
	default:
		log.Println("Unknown language. Falling back to EN\n")
		setupENConfig()
	}
}

// directories, default config options
func setupLanguageNeutralConfig() {

	config.sourceDirName = GetUserHomeDirectory()
	config.targetDirName = GetUserHomeDirectory()

	config.numerate = true
	config.pageNumberPrefix = ""
	config.chapterPageSeparator = " - "
	config.chapterPrefix = ""

	config.headerText = ""

	config.evenify = false

	config.blankPageText = ""

	config.concatenate = false

	config.resultingPDFFileName = ""

}

func setupDEConfig() {
	config.pageNumberPrefix = "Seite"
	config.chapterPrefix = "Kapitel"

	config.blankPageText = "Diese Seite bleibt absichtlich frei"

}

func setupENConfig() {
	config.pageNumberPrefix = "Page"
	config.chapterPrefix = "Chapter"

	config.blankPageText = "Page intentionally left blank"
}

func checkPreferredLanguage() string {

	var userPrefs = []language.Tag{
		language.Make("de"), // German
		//language.Make("fr"),  // French
	}

	var serverLangs = []language.Tag{
		language.AmericanEnglish, // en-US fallback
		language.German,          // de
	}

	var matcher = language.NewMatcher(serverLangs)

	tag, _, _ := matcher.Match(userPrefs...)

	fmt.Printf("best match: %s (%s)\n",
		display.English.Tags().Name(tag),
		display.Self.Name(tag))

	return display.English.Tags().Name(tag)
}

// a single PDF file
type singleFileToProcess struct {
	directory        string
	filename         string
	origPageCount    int
	hasBeenEvenified bool
}

// our "to-do" list with all files that need to be processed
type processingState struct {
	totalPageCount        int
	pagesAlreadyProcessed int

	filesToProcess []singleFileToProcess
}
