package translator

import "strings"

type (
	Translator interface {
		TranslateEN(key string) string
		Translate(key string, lang Lang) string
	}

	Lang string
)

const (
	EN = "en"
	FA = "fa"
)

func GetLanguage(lang string) Lang {
	switch strings.ToLower(lang) {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
