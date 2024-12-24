// Package itfilez provides internal test fixtures for the "filez" package.
package itfilez

import (
	"embed"
)

var (
	// AssetsEmbed provides an embedded FS for tests.
	//
	//go:embed assets
	AssetsEmbed embed.FS
)
