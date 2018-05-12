package utils

import (
	"os"
	"log"
)


func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
}
