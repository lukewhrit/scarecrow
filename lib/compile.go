/*
 * Copyright © 2020-2022 Luke Whritenour <lukewhrit@proton.me>

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 *  http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lib

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

// Document is a Scarecrow document file
type Document struct {
	FileInfo os.FileInfo
	Content  []byte
	Layout   []byte
	FullPath string
}

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		extension.Footnote,
	),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

func (d *Document) compileLayout() []byte {
	re := regexp.MustCompile("(?i){{ *body *}}")
	return re.ReplaceAll(d.Layout, d.Content)
}

func (d *Document) compileMarkdown() ([]byte, error) {
	var buf bytes.Buffer
	err := goldmark.Convert(d.Content, &buf)
	return buf.Bytes(), err
}

func (d *Document) writeFile(dir, subDir, output string) (err error) {
	content := d.compileLayout()
	content, err = MinifyHTML(content)

	outputFolder := filepath.Join(dir, output, subDir)
	outputFile := filepath.Join(
		outputFolder,
		strings.ReplaceAll(d.FileInfo.Name(), ".md", ".html"),
	)

	err = os.MkdirAll(outputFolder, os.ModePerm)
	return ioutil.WriteFile(outputFile, []byte(content), 0600)
}

// Compile turns a markdown document into a fully-formed HTML file using a layout
func (d *Document) Compile(dir, output string) (err error) {
	d.Content, err = ioutil.ReadFile(d.FullPath)
	d.Content, err = d.compileMarkdown()
	d.Content, err = CompileTemplate(d.FileInfo.Name(), d.Content)

	path, err := filepath.Rel(dir, d.FullPath)
	folder := strings.Split(path, string(filepath.Separator))[0]
	subDir := folder

	// Pages get sent to the root of the output directory instead of a
	// sub-directory, like other files.
	if folder == "pages" {
		subDir = ""
	}

	return d.writeFile(dir, subDir, output)
}
