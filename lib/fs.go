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
