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
	"bytes"
	"html/template"
	"log"
)

// CompileTemplate compiles an HTML template
func CompileTemplate(name string, raw []byte) []byte {
	tmpl, err := template.New(name).Parse(string(raw))
	Handle(err)

	var compiledTemplate bytes.Buffer
	if err := tmpl.ExecuteTemplate(&compiledTemplate, name, raw); err != nil {
		log.Fatalf(err.Error())
	}

	return compiledTemplate.Bytes()
}
