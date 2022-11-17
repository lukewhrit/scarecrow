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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/karrick/godirwalk"
)

// MoveAssets loads all asset files from directory and moves them into the
// `outputDir`.
func MoveAssets(baseDir, outputDir string) (err error) {
	dir := filepath.Join(baseDir, "assets")

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return godirwalk.Walk(dir, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) (err error) {
			if de.IsDir() {
				return
			}

			data, err := ioutil.ReadFile(path)
			info, err := os.Stat(path)

			output := filepath.Join(outputDir, info.Name())
			err = os.MkdirAll(outputDir, os.ModePerm)

			return ioutil.WriteFile(output, data, os.ModePerm)
		},
		Unsorted: true,
	})
}
