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
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/lukewhrit/scarecrow/util"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)

var clean bool
var output string

var makeCmd = &cobra.Command{
	Use:   "make <dir>",
	Short: "Compile a Scarecrow project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(args[0])
		util.Handle(err)

		files := []string{}
		// Loop over files in provided path and add to array
		if err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			files = append(files, path)

			return err
		}); err != nil {
			log.Fatalf(err.Error())
		}

		fileContents := map[string][]byte{}

		for _, file := range files {
			info, err := os.Stat(file)
			util.Handle(err)

			// Make sure we don't include directories
			if !info.IsDir() {
				// Only use files with the correct extension/name
				if strings.HasSuffix(file, ".md") || info.Name() == "layout.html" {
					content, err := ioutil.ReadFile(file)
					util.Handle(err)

					// If file is layout don't run blackfriday on it
					if info.Name() == "layout.html" {
						fileContents[info.Name()] = content
					} else {
						fileContents[info.Name()] = blackfriday.Run(content)
					}

					if strings.HasSuffix(file, ".md") {
						pathItems := strings.Split(file, string(filepath.Separator))
						// Get second item in array
						// This value will most likely be the bottom-level folder the files are in
						folder := pathItems[len(pathItems)-2]
						subdir := ""

						if folder == "pages" {
							subdir = ""
						} else if folder == "posts" {
							subdir = "posts"
						}

						if err := util.WriteFile(fileContents, dir, subdir, info.Name()); err != nil {
							log.Fatal(err.Error())
						}
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
