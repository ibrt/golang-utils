// Package hashz provides various utilities calculating hashes.
package hashz

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"hash/fnv"

	"github.com/ibrt/golang-utils/errorz"
)

// MustHashMD5 calculates a MD5 hash and returns it as a hex-encoded string.
func MustHashMD5(buf []byte) string {
	return MustHash(md5.New(), buf)
}

// MustHashSHA256 calculates an SHA-256 hash  and returns it as a hex-encoded string.
func MustHashSHA256(buf []byte) string {
	return MustHash(sha256.New(), buf)
}

// MustHashSHA1 calculates an SHA-1 hash and returns it as a hex-encoded string.
func MustHashSHA1(buf []byte) string {
	return MustHash(sha1.New(), buf)
}

// MustHashFNV1128 calculates a 128-bit FNV-1 hash and returns it as a hex-encoded string.
func MustHashFNV1128(buf []byte) string {
	return MustHash(fnv.New128(), buf)
}

// MustHash completes the hash by writing the buffer, summing it, and returning it as a hex-encoded string.
func MustHash(h hash.Hash, buf []byte) string {
	_, err := h.Write(buf)
	errorz.MaybeMustWrap(err)
	return hex.EncodeToString(h.Sum(nil))
}
