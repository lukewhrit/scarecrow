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
	"log"
	"path/filepath"
)

var pathSeperator = string(filepath.Separator)

// Handle handles errors
func Handle(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// HasExt gets and validates a string's file extension
func HasExt(str, ext string) bool {
	if filepath.Ext(str) == fmt.Sprintf(".%s", ext) {
		return true
	}

	return false
}

// Contains checks if a slice contains a string
func Contains(slice []string, str string) bool {
	for _, a := range slice {
		if a == str {
			return true
		}
	}

	return false
}
