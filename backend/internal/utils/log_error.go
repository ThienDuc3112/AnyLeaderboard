package utils

import "log"

func LogError(handlerName string, e error) {
	if e != nil {
		log.Printf("\n========== %s error ==========\n%v\n========================================\n\n", handlerName, e)
	}
}
