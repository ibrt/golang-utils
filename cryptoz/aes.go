package cryptoz

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"strings"

	"github.com/ibrt/golang-utils/errorz"
)

// AES constants.
const (
	CiphertextPrefix = "enc:"
	KeyLen           = 32
	EncodedKeyLen    = 64
)

// MustEncryptAES encrypts the given plaintext using symmetric AES, GCM mode.
// The key must consist of 32 bytes, hex-encoded (i.e. 64 characters).
// The ciphertext will be base64-encoded, prefixed with "enc:".
func MustEncryptAES(key, plaintext string) string {
	errorz.Assertf(!strings.HasPrefix(plaintext, CiphertextPrefix), "already encrypted")

	rawKey, err := hex.DecodeString(key)
	errorz.MaybeMustWrap(err)
	errorz.Assertf(len(rawKey) == 32, "invalid key length")

	c, err := aes.NewCipher(rawKey)
	errorz.MaybeMustWrap(err)

	gcm, err := cipher.NewGCM(c)
	errorz.MaybeMustWrap(err)

	nonceBuf := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonceBuf)
	errorz.MaybeMustWrap(err)

	buf := gcm.Seal(nonceBuf, nonceBuf, []byte(plaintext), nil)
	return CiphertextPrefix + base64.StdEncoding.EncodeToString(buf)
}

// MustDecryptAES decrypts a ciphertext encrypted by MustEncryptAES.
// The key must consist of 32 bytes, hex-encoded (i.e. 64 characters).
// The ciphertext must be base64-encoded, prefixed with "enc:".
func MustDecryptAES(key, ciphertext string) string {
	errorz.Assertf(strings.HasPrefix(ciphertext, CiphertextPrefix), "not encrypted")

	rawKey, err := hex.DecodeString(key)
	errorz.MaybeMustWrap(err)
	errorz.Assertf(len(rawKey) == 32, "invalid key length")

	c, err := aes.NewCipher(rawKey)
	errorz.MaybeMustWrap(err)

	gcm, err := cipher.NewGCM(c)
	errorz.MaybeMustWrap(err)

	buf, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(ciphertext, CiphertextPrefix))
	errorz.MaybeMustWrap(err)
	errorz.Assertf(len(buf) >= gcm.NonceSize(), "invalid ciphertext")

	plaintext, err := gcm.Open(nil, buf[:gcm.NonceSize()], buf[gcm.NonceSize():], nil)
	errorz.MaybeMustWrap(err)
	return string(plaintext)
}

// MustGenerateRandomAESKey generates a random AES key.
func MustGenerateRandomAESKey() string {
	buf := make([]byte, KeyLen)
	_, err := rand.Read(buf)
	errorz.MaybeMustWrap(err)
	return hex.EncodeToString(buf)
}
