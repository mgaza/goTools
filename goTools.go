package goTools

import "log"

func CheckErrorFatal(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
