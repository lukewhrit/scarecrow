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
	"path/filepath"

	"github.com/karrick/godirwalk"
)

var (
	files []string
	info  os.FileInfo
	de    os.FileInfo
	data  []byte
)

// MoveAssets loads all asset files from directory and moves them into the `dist/` directory
func MoveAssets(baseDir, outputDir string) (err error) {
	dir := filepath.Join(baseDir, "assets")

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	if err := godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) (err error) {
			files = append(files, path)

			return nil
		},
		Unsorted: true,
	}); err != nil {
		return err
	}

	for _, path := range files {
		de, err = os.Stat(path)

		if de.IsDir() {
			continue
		}

		data, err = ioutil.ReadFile(path)
		info, err = os.Stat(path)

		output := filepath.Join(outputDir, info.Name())
		fmt.Println(output)
		err = os.MkdirAll(outputDir, os.ModePerm)

		return ioutil.WriteFile(output, data, os.ModePerm)
	}

	return nil
}
