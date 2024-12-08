package domain

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
	"os"
	"strings"
)

var supportedLanguages = []language.Tag{
	language.German,
	language.English,
}

var matcher = language.NewMatcher(supportedLanguages)

func DetectSystemLanguage() language.Tag {
	sysLocale := os.Getenv("LANG")
	if sysLocale == "" {
		sysLocale = os.Getenv("LC_ALL")
	}

	tag, err := language.Parse(strings.Split(sysLocale, ".")[0])
	if err != nil {
		return language.English // fallback
	}

	tag, _, _ = matcher.Match(tag)
	return tag
}

func GetLanguageName(tag language.Tag) string {
	namer := display.Tags(tag) // This is the correct way
	return namer.Name(tag)
}

func GetLanguageNameInEnglish(tag language.Tag) string {
	namer := display.English.Tags() // This is the correct way
	return namer.Name(tag)
}

func ValidateLanguage(lang string) (language.Tag, error) {
	tag, err := language.Parse(lang)
	if err != nil {
		return language.English, err
	}

	tag, _, _ = matcher.Match(tag)
	return tag, nil
}

func ListAvailableLanguages() [][]string {
	languages := [][]string{}
	for _, tag := range supportedLanguages {
		// Get language code (two-letter code)
		langCode := tag.String()

		// Get language name in the language itself
		nameInLanguage := GetLanguageName(tag)

		// Get language name in English
		nameInEnglish := GetLanguageNameInEnglish(tag)

		languages = append(languages, []string{
			langCode,       // Language code
			nameInLanguage, // Name in native language
			nameInEnglish,  // Name in English
		})
	}
	return languages
}

func PrintLanguages() {
	languages := ListAvailableLanguages()
	for _, lang := range languages {
		// Each `lang` is a []string with the format [langCode, nameInLanguage, nameInEnglish]
		fmt.Printf("Code: %s, Native: %s, English: %s\n", lang[0], lang[1], lang[2])
	}

}
