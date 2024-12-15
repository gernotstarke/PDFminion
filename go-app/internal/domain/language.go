package domain

import (
	"fmt"
	"github.com/Xuanwo/go-locale"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"os"
	"strings"
)

var supportedLanguages = []language.Tag{
	language.German,
	language.English,
	language.French,
}

var matcher = language.NewMatcher(supportedLanguages)

// DetectSystemLanguage detects the system language by checking environment variables and falling back to locale.Detect().
func DetectSystemLanguage() language.Tag {
	// First, check LANG and LC_ALL environment variables
	sysLocale := os.Getenv("LANG")
	if sysLocale == "" {
		sysLocale = os.Getenv("LC_ALL")
	}

	if sysLocale != "" {
		tag, err := language.Parse(strings.Split(sysLocale, ".")[0])
		if err == nil {
			tag, _, _ = matcher.Match(tag)
			base, _ := tag.Base()               // Normalize to base language
			return language.Make(base.String()) // Convert Base to Tag
		}
	}

	// Fallback to locale.Detect()
	tag, err := locale.Detect()
	if err == nil {
		tag, _, _ = matcher.Match(tag)
		base, _ := tag.Base()               // Normalize to base language
		return language.Make(base.String()) // Convert Base to Tag
	}

	// Fallback to default language
	return language.English
}

// GetLanguageName returns the name of the language in its native form.
func GetLanguageName(tag language.Tag) (code, nameInOriginal, nameInEnglish string) {
	base, _ := tag.Base()                            // Get the base language tag (e.g., "de" from "de-u-rg-dezzzz")
	code = base.String()                             // Simplified code (e.g., "de", "en", "fr")
	nameInOriginal = display.Self.Name(tag)          // Name in its native language
	nameInEnglish = display.English.Tags().Name(tag) // Name in English
	return
}

func ValidateLanguage(lang string) (language.Tag, error) {
	tag, err := language.Parse(lang)
	if err != nil {
		return language.English, err
	}

	tag, _, _ = matcher.Match(tag)
	return tag, nil
}

// ListAvailableLanguages returns a list of supported languages with their codes and names.
func ListAvailableLanguages() [][]string {
	supportedLanguages := []language.Tag{
		language.English,
		language.German,
		language.French,
	}

	var languages [][]string
	for _, tag := range supportedLanguages {
		langCode := tag.String()                          // e.g., "en", "de", "fr"
		nameInEnglish := display.English.Tags().Name(tag) // Name in English
		nameInOriginal := display.Self.Name(tag)          // Name in its native language
		languages = append(languages, []string{langCode, nameInOriginal, nameInEnglish})
	}
	return languages
}

func PrintLanguages() {
	currentLanguage := DetectSystemLanguage()
	currentCode, currentNameInOriginal, currentNameInEnglish := GetLanguageName(currentLanguage)

	fmt.Println("Supported Languages:")
	languages := ListAvailableLanguages()
	for _, lang := range languages {
		fmt.Printf("Code: %s, %s (%s)\n", lang[0], lang[1], lang[2])
	}

	// Print the current language with its details
	fmt.Printf("\nYour current language is  %s (%s, %s)\n",
		currentCode, currentNameInOriginal, currentNameInEnglish)
}
