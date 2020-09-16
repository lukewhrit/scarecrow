/*
 * Copyright © 2020 Luke Whrit <lukewhrit@gmail.com>

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
