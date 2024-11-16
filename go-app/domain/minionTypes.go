package domain

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"log"
	"os"
	"os/user"
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

	pageNumberPrefix     string
	chapterPageSeparator string
	chapterPrefix        string

	// evenify: make every PDF have an EVEN page count,
	// by adding a single blank page to those PDFs with
	// originally odd number of pages
	evenify bool

	// this String gets stamped onto pages added during evenification
	blankPageText string "Diese Seite bleibt absichtlich frei"
} // MinionConfiguration

// SetupConfiguration initializes the configuration settings
func SetupConfiguration() {

	setupLanguageNeutralConfig()

	switch lang := checkPreferredLanguage(); lang {
	case "German":
		log.Printf("Deutsch als Sprache identifiziert.\n")
		setupDEConfig()
	case "English":
		log.Printf("English identified as user language.")
		setupENConfig()
	default:
		log.Printf("Unknown language. Falling back to EN\n")
		setupENConfig()
	}
}

// directories, default config options
func setupLanguageNeutralConfig() {

	config.sourceDirName = GetUserHomeDirectory()
	config.targetDirName = GetUserHomeDirectory()

	config.pageNumberPrefix = ""
	config.chapterPageSeparator = " - "
	config.chapterPrefix = ""

	config.evenify = false

	config.blankPageText = ""

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

// GetUserHomeDirectory returns the current user's home directory.
func GetUserHomeDirectory() string {
	usr, err := user.Current()
	if err != nil {
		return os.Getenv("HOME") // Fallback to HOME environment variable
	}
	return usr.HomeDir
}
