package utils

import (
	"fmt"
	"log"
	"strings"
)

func LogError(handlerName string, e error) {
	if e != nil {
		topBar := fmt.Sprintf("========== %s error ==========", handlerName)
		bottomBar := strings.Map(func(r rune) rune { return '=' }, topBar)
		log.Printf("\n%s\n%v\n%s\n\n", topBar, e, bottomBar)
	}
}
