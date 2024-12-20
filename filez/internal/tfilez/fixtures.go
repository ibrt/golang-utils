// Package tfilez provides test fixtures for the "filez" package.
package tfilez

import (
	"embed"
)

var (
	// FixturesEmbed provides an embedded FS for tests.
	//
	//go:embed fixtures
	FixturesEmbed embed.FS
)
