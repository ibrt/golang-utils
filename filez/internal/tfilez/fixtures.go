package tfilez

import (
	"embed"
)

// FixturesEmbed provides an embedded FS for tests.
//
//go:embed fixtures
var FixturesEmbed embed.FS
