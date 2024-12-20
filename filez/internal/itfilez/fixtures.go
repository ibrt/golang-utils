// Package itfilez provides internal test fixtures for the "filez" package.
package itfilez

import (
	"embed"
)

var (
	// FixturesEmbed provides an embedded FS for tests.
	//
	//go:embed fixtures
	FixturesEmbed embed.FS
)
