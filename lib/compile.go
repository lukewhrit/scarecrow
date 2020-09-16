package lib

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// Document is a Scarecrow document file
type Document struct {
	FileInfo os.FileInfo
	Content  []byte
	Layout   []byte
	FullPath string
}

// Compile turns a markdown document into a fully-formed HTML file using a layout
func (d *Document) Compile(dir string) (err error) {
	d.Content, err = ioutil.ReadFile(d.FullPath)
	d.Content, err = MinifyHTML(blackfriday.Run(d.Content))

	path, err := filepath.Rel(dir, d.FullPath)
	folder := strings.Split(path, string(filepath.Separator))[0]
	subDir := folder

	if folder == "pages" {
		subDir = ""
	}

	return d.WriteFile(dir, subDir)
}
