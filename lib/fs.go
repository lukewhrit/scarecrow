package lib

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// WriteFile writes a file to a `dir`
func (d *Document) WriteFile(dir, subDir string) error {
	content := strings.ReplaceAll(string(d.Layout), "<scarecrow-body />", string(d.Content))
	outputFile := fmt.Sprintf("%s/dist/%s/%s", dir, subDir,
		strings.ReplaceAll(d.FileInfo.Name(), ".md", ".html"))

	return ioutil.WriteFile(outputFile, []byte(content), 0600)
}
