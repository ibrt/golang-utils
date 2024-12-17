package cryptoz

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/ibrt/golang-utils/errorz"
)

const (
	privateKeyHeader = "PRIVATE KEY"
	publicKeyHeader  = "PUBLIC KEY"
)

// PEMToRSAPrivateKey tries to find and parse a [*rsa.PrivateKey] in the given PEM file.
func PEMToRSAPrivateKey(buf []byte) (*rsa.PrivateKey, error) {
	for {
		var block *pem.Block
		block, buf = pem.Decode(buf)

		if block == nil {
			return nil, errorz.Errorf("invalid PEM file or RSA private key not found")
		}

		if block.Type != privateKeyHeader {
			continue
		}

		rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			continue
		}

		if rsaKey, ok := rawKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
	}
}

// MustPEMToRSAPrivateKey is like [PEMToRSAPrivateKey] but panics on error.
func MustPEMToRSAPrivateKey(buf []byte) *rsa.PrivateKey {
	key, err := PEMToRSAPrivateKey(buf)
	errorz.MaybeMustWrap(err)
	return key
}

// PEMToRSAPublicKey tries to find and parse a [*rsa.PublicKey] in the given PEM file.
func PEMToRSAPublicKey(buf []byte) (*rsa.PublicKey, error) {
	for {
		var block *pem.Block
		block, buf = pem.Decode(buf)

		if block == nil {
			return nil, errorz.Errorf("invalid PEM file or RSA public key not found")
		}

		if block.Type != publicKeyHeader {
			continue
		}

		rawKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			continue
		}

		if rsaKey, ok := rawKey.(*rsa.PublicKey); ok {
			return rsaKey, nil
		}
	}
}

// MustPEMToRSAPublicKey is like [PEMToRSAPublicKey] but panics on error.
func MustPEMToRSAPublicKey(buf []byte) *rsa.PublicKey {
	key, err := PEMToRSAPublicKey(buf)
	errorz.MaybeMustWrap(err)
	return key
}

// RSAPrivateKeyToPEM encodes a [*rsa.PrivateKey] to PEM format.
func RSAPrivateKeyToPEM(key *rsa.PrivateKey) []byte {
	buf, err := x509.MarshalPKCS8PrivateKey(key)
	errorz.MaybeMustWrap(err) // never triggers because we already checked the key type

	return pem.EncodeToMemory(&pem.Block{
		Type:  privateKeyHeader,
		Bytes: buf,
	})
}

// RSAPublicKeyToPEM encodes a [*rsa.PublicKey] to PEM format.
func RSAPublicKeyToPEM(key *rsa.PublicKey) []byte {
	buf, err := x509.MarshalPKIXPublicKey(key)
	errorz.MaybeMustWrap(err) // never triggers because we already checked the key type

	return pem.EncodeToMemory(&pem.Block{
		Type:  publicKeyHeader,
		Bytes: buf,
	})
}
