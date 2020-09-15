package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Handle handles errors
func Handle(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// WriteFile writes a file to a `dir`
func WriteFile(fileContents map[string][]byte, dir, subdir, name string) error {
	content := strings.Replace(
		string(fileContents["layout.html"]),
		"<scarecrow-body>",
		string(fileContents[name]),
		1)

	outputFile := fmt.Sprintf("%s/dist/%s/%s", dir,
		subdir,
		strings.ReplaceAll(name, ".md", ".html"))

	return ioutil.WriteFile(outputFile, []byte(content), 0600)
}
