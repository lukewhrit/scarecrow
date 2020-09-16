package lib

import (
	"fmt"
	"log"
	"path/filepath"
)

// Handle handles errors
func Handle(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// HasExt gets and validates a string's file extension
func HasExt(str, ext string) bool {
	if filepath.Ext(str) == fmt.Sprintf(".%s", ext) {
		return true
	}

	return false
}
