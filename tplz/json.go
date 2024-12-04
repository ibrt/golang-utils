package tplz

import (
	"bytes"
	"encoding/json"
	ttpl "text/template"

	"github.com/ibrt/golang-utils/errorz"
)

// ExecuteJSON executes a text template, formatting the result as JSON code.
func ExecuteJSON(template *ttpl.Template, prefix, indent string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := template.Execute(buf, data); err != nil {
		return nil, errorz.Wrap(err)
	}

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, buf.Bytes(), prefix, indent); err != nil {
		return nil, errorz.Wrap(err)
	}

	return dst.Bytes(), nil
}

// MustExecuteJSON executes a text template, formatting the result as JSON code, panics on error.
func MustExecuteJSON(template *ttpl.Template, prefix, indent string, data any) []byte {
	buf, err := ExecuteJSON(template, prefix, indent, data)
	errorz.MaybeMustWrap(err)
	return buf
}

// ParseAndExecuteJSON parses and executes a text template, formatting the result as JSON code.
func ParseAndExecuteJSON(template, prefix, indent string, data any) ([]byte, error) {
	parsedTemplate, err := ttpl.New("").Parse(template)
	if err != nil {
		return nil, errorz.Wrap(err)
	}

	return ExecuteJSON(parsedTemplate, prefix, indent, data)
}

// MustParseAndExecuteJSON is like ParseAndExecuteJSON but panics on error.
func MustParseAndExecuteJSON(template, prefix, indent string, data any) []byte {
	buf, err := ParseAndExecuteJSON(template, prefix, indent, data)
	errorz.MaybeMustWrap(err)
	return buf
}
