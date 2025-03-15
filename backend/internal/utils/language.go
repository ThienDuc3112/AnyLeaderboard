package utils

import "github.com/abadojack/whatlanggo"

func mapLangToPostgres(langCode string) string {
	mapping := map[string]string{
		"eng": "english",
		"fra": "french",
		"deu": "german",
		"spa": "spanish",
		"ita": "italian",
		"nld": "dutch",
		"rus": "russian",
		"por": "portuguese",
		"dan": "danish",
		"fin": "finnish",
		"hun": "hungarian",
		"nor": "norwegian",
		"ron": "romanian",
		"swe": "swedish",
		"tur": "turkish",
		"cat": "catalan",
		"ell": "greek",
		"ind": "indonesian",
		"gle": "irish",
		"lit": "lithuanian",
		"nep": "nepali",
		"srp": "serbian",
		"slk": "slovak",
		"slv": "slovenian",
		"tam": "tamil",
		"tha": "thai",
	}

	// Return the mapped PostgreSQL `regconfig`, or default to 'simple'
	if pgLang, exists := mapping[langCode]; exists {
		return pgLang
	}
	return "simple" // Default for unsupported languages
}

func DetectLanguage(text string) string {
	info := whatlanggo.Detect(text)
	return mapLangToPostgres(whatlanggo.LangToString(info.Lang))
}
