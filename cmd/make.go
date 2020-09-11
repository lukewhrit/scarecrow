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

		for _, file := range files {
			info, err := os.Stat(file)

			if err != nil {
				log.Fatalf(err.Error())
			}

			// Make sure we don't include directories
			if !info.IsDir() {
				content, err := ioutil.ReadFile(file)

				if err != nil {
					log.Fatalf(err.Error())
				}

				if strings.HasSuffix(file, ".md") {
					fmt.Printf("%s ---\n", info.Name())
					fmt.Println(string(blackfriday.Run(content)))
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
}
