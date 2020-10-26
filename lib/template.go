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
	"github.com/cbroglie/mustache"
)

var hash = map[string]interface{}{
	"c": "world",
	"repo": []map[string]string{
		{"name": "resque"},
		{"name": "hub"},
		{"name": "rip"},
	},
}

// CompileTemplate compiles an HTML template
func CompileTemplate(name string, raw []byte) ([]byte, error) {
	rawString := string(raw)
	data, err := mustache.Render(rawString, hash)

	return []byte(data), err
}
