package domain

import (
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
