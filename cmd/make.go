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
	dest               string
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

		assetsDest, err := filepath.Abs(filepath.Join(dir, dest))
		lib.Handle(err)

		if err := lib.MoveAssets(dir, assetsDest); err != nil {
			log.Fatal(err.Error())
		}

		layout, err := ioutil.ReadFile(fmt.Sprintf("%s%slayout.html",
			dir, string(filepath.Separator)))
		lib.Handle(err)

		// Walk project directory and compile files
		lib.Handle(godirwalk.Walk(dir, &godirwalk.Options{
			Callback: func(path string, de *godirwalk.Dirent) error {
				if !de.IsDir() || de.Name() != layoutFileName {
					relPath, err := filepath.Rel(dir, path)
					tlDir := strings.Split(relPath, string(filepath.Separator))[0]

					if err != nil {
						return err
					}

					if !lib.Contains(allowedDirectories, tlDir) {
						return nil
					}

					if filepath.Ext(de.Name()) == ".md" {
						fileInfo, err := os.Stat(path)

						if err != nil {
							return err
						}

						doc := &lib.Document{
							FileInfo: fileInfo,
							Content:  []byte{},
							Layout:   layout,
							FullPath: path,
						}

						return doc.Compile(dir, dest)
					}
				}

				return nil
			},
			Unsorted: true,
		}))
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	makeCmd.Flags().BoolVarP(&clean, "clean", "c", true, "cleanup directory before saving new output")
	makeCmd.Flags().StringVarP(&dest, "dest", "o", "./dist", "send output to a custom directory")
}
