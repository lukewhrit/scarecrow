package util

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

func MinifyHTML(content []byte) ([]byte, error) {
	m := minify.New()

	m.AddFunc("text/html", html.Minify)

	return m.Bytes("text/html", content)
}
