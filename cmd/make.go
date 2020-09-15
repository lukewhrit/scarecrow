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

		// Loop over files in provided path and add to array
		files := []string{}
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
				// Only use files with the correct extension
				if strings.HasSuffix(info.Name(), ".md") || strings.HasSuffix(info.Name(), ".html") {
					content, err := ioutil.ReadFile(file)
					util.Handle(err)

					// If file is layout don't run blackfriday on it
					if strings.HasSuffix(info.Name(), ".html") {
						minifiedContent, err := util.MinifyHTML(content)
						util.Handle(err)

						fileContents[info.Name()] = minifiedContent
					} else {
						minifiedContent, err := util.MinifyHTML(blackfriday.Run(content))
						util.Handle(err)

						fileContents[info.Name()] = minifiedContent
					}

					if strings.HasSuffix(file, ".md") {
						path, err := filepath.Rel(dir, file)
						util.Handle(err)

						// Create a relative path between the * and *
						folder := strings.Split(path, string(filepath.Separator))[0]
						subDir := ""

						switch folder {
						case "pages":
							subDir = ""
						default:
							subDir = folder
						}

						if err := util.WriteFile(fileContents, dir, subDir, info.Name()); err != nil {
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
