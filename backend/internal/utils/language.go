package utils

import "github.com/abadojack/whatlanggo"

func DetectLanguage(text string) string {
	info := whatlanggo.Detect(text)
	return whatlanggo.LangToString(info.Lang)
}
