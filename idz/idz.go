package idz

import (
	"github.com/google/uuid"
)

// MustNewRandomUUID generates a new random ID.
func MustNewRandomUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// IsValidUUID returns true if the given ID is valid.
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
