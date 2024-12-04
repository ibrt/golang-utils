package tplz

import (
	"bytes"
	"go/format"
	ttpl "text/template"

	"github.com/ibrt/golang-utils/errorz"
)

// ExecuteGo executes a text template, formatting the result as Go code.
func ExecuteGo(template *ttpl.Template, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, errorz.Wrap(err)
	}

	return format.Source(buf.Bytes())
}

// MustExecuteGo is like ExecuteGo but panics on error.
func MustExecuteGo(template *ttpl.Template, data any) []byte {
	buf, err := ExecuteGo(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}

// ParseAndExecuteGo parses and executes a text template, formatting the result as Go code.
func ParseAndExecuteGo(template string, data any) ([]byte, error) {
	parsedTemplate, err := ttpl.New("").Parse(template)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	return ExecuteGo(parsedTemplate, data)
}

// MustParseAndExecuteGo is like ParseAndExecuteGo but panics on error.
func MustParseAndExecuteGo(template string, data any) []byte {
	buf, err := ParseAndExecuteGo(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}
