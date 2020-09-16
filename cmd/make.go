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
	"log"
	"os"
	"path/filepath"

	"github.com/lukewhrit/scarecrow/lib"
	"github.com/spf13/cobra"
)

var clean bool
var output string

var makeCmd = &cobra.Command{
	Use:   "make <dir>",
	Short: "Compile a Scarecrow project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(args[0]) // Get project base directory
		lib.Handle(err)

		// Walk project directory and add files to slice
		files := []string{}
		lib.Handle(filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			files = append(files, path)
			return err
		}))

		// Loop over every file in directory and compile
		for _, file := range files {
			info, err := os.Stat(file)
			lib.Handle(err)

			// Make sure we don't try to build directories or the layout file
			if !info.IsDir() || info.Name() != "layout.html" {
				// Don't try to compile any files with an invalid extension, such as CSS or JS files.
				if lib.HasExt(info.Name(), "md") {
					doc := &lib.Document{
						FileInfo: info,
						Content:  []byte{},                     // Should be empty, `Compile` method will fill it in.
						Layout:   []byte("<scarecrow-body />"), // @todo Get layout file content and place here
						FullPath: file,
					}

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
