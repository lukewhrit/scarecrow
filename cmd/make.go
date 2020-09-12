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

	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)

var dryRun bool
var clean bool
var output string

var makeCmd = &cobra.Command{
	Use:   "make <dir>",
	Short: "Compile a Scarecrow project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := filepath.Abs(args[0])

		if err != nil {
			log.Fatalf(err.Error())
		}

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

			if err != nil {
				log.Fatalf(err.Error())
			}

			// Make sure we don't include directories
			if !info.IsDir() {
				// Only use files with the correct extension/name
				if strings.HasSuffix(file, ".md") || info.Name() == "layout.html" {
					content, err := ioutil.ReadFile(file)

					if err != nil {
						log.Fatalf(err.Error())
					}

					// If file is layout don't run blackfriday on it
					if info.Name() == "layout.html" {
						fileContents[info.Name()] = content
					} else {
						fileContents[info.Name()] = blackfriday.Run(content)
					}

					if strings.HasSuffix(file, ".md") {
						if dryRun {
							fmt.Printf("----- %s \n", info.Name())
							fmt.Println(string(fileContents[info.Name()]))
						} else {
							content := strings.Replace(
								string(fileContents["layout.html"]),
								"<scarecrow-body />",
								string(fileContents[info.Name()]),
								1)

							err := ioutil.WriteFile(
								strings.Replace(file, ".md", ".html", 1),
								[]byte(content),
								0600)

							if err != nil {
								log.Fatalf(err.Error())
							}
						}
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	makeCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "only output to stdout")
	makeCmd.Flags().BoolVarP(&clean, "clean", "c", true, "cleanup directory before saving new output")
	makeCmd.Flags().StringVarP(&output, "output", "o", "./dist", "send output to a custom directory")
}
