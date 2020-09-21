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

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/lukewhrit/scarecrow/lib"
	"github.com/spf13/cobra"
)

var (
	clean              bool
	output             string
	files              []string
	allowedDirectories = []string{"pages", "posts"}
)

const layoutFileName = "layout.html"

var makeCmd = &cobra.Command{
	Use:   "make <dir>",
	Short: "Compile a Scarecrow project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(args[0]) // Get project base directory
		lib.Handle(err)

		// Walk project directory and add files to slice
		lib.Handle(godirwalk.Walk(dir, &godirwalk.Options{
			Callback: func(path string, de *godirwalk.Dirent) error {
				files = append(files, path)
				return nil
			},
			Unsorted: true,
		}))

		layout, err := ioutil.ReadFile(fmt.Sprintf("%s%slayout.html",
			dir, string(filepath.Separator)))
		lib.Handle(err)

		// Loop over every file in directory and compile
		for _, filePath := range files {
			info, err := os.Stat(filePath)
			lib.Handle(err)

			// Make sure we don't try to build directories or the layout file
			if !info.IsDir() || info.Name() != layoutFileName {
				// Don't try to compile any files with an invalid extension, such as CSS or JS files.
				if lib.HasExt(info.Name(), "md") {
					path, err := filepath.Rel(dir, filePath)
					lib.Handle(err)

					// Only compile files within the allowed directories
					if !lib.Contains(allowedDirectories,
						strings.Split(path, string(filepath.Separator))[0]) {
						continue
					}

					doc := &lib.Document{
						FileInfo: info,
						Content:  []byte{}, // Should be empty, `Compile` method will fill it in.
						Layout:   layout,
						FullPath: filePath,
					}

					// Compile the document
					if err := doc.Compile(dir); err != nil {
						log.Fatal(err.Error())
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	makeCmd.Flags().BoolVarP(&clean, "clean", "c", true, "cleanup directory before saving new output")
	makeCmd.Flags().StringVarP(&output, "output", "o", "./dist", "send output to a custom directory")
}
