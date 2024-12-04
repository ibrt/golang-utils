package tplz

import (
	"bytes"
	htpl "html/template"

	"github.com/yosssi/gohtml"

	"github.com/ibrt/golang-utils/errorz"
)

// ExecuteHTML executes a HTML template.
func ExecuteHTML(template *htpl.Template, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(gohtml.NewWriter(buf), data); err != nil {
		return nil, errorz.Wrap(err)
	}

	return buf.Bytes(), nil
}

// MustExecuteHTML is like ExecuteHTML but panics on error.
func MustExecuteHTML(template *htpl.Template, data any) []byte {
	buf, err := ExecuteHTML(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}

// ParseAndExecuteHTML parses and executes a HTML template.
func ParseAndExecuteHTML(template string, data any) ([]byte, error) {
	parsedTemplate, err := htpl.New("").Parse(template)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	return ExecuteHTML(parsedTemplate, data)
}

// MustParseAndExecuteHTML is like ParseAndExecuteHTML but panics on error.
func MustParseAndExecuteHTML(template string, data any) []byte {
	buf, err := ParseAndExecuteHTML(template, data)
	errorz.MaybeMustWrap(err)
	return buf
}
