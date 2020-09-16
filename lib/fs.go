package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// WriteFile writes a file to a `dir`
func (d *Document) WriteFile(dir, subDir string) (err error) {
	content := strings.ReplaceAll(string(d.Layout), "<scarecrow-body />", string(d.Content))

	outputFolder := fmt.Sprintf("%s/dist/%s", dir, subDir)
	outputFile := fmt.Sprintf("%s/%s", outputFolder,
		strings.ReplaceAll(d.FileInfo.Name(), ".md", ".html"))

	err = os.MkdirAll(outputFolder, os.ModePerm)
	return ioutil.WriteFile(outputFile, []byte(content), 0600)
}
