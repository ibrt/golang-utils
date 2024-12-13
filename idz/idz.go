// Package idz provides various utilities for generating IDs and random values.
package idz

import (
	"github.com/google/uuid"
)

// MustNewRandomUUID generates a new random UUID v4.
func MustNewRandomUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// IsValidUUID returns true if the given UUID is well-formed.
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
