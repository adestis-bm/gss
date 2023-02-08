package gss

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type SecureKey struct {
	phrase []byte
}

type SecureString struct {
	nonce  []byte
	sealed []byte
}

// New prepares a secure string with a new key.
func New(u string) (*SecureString, *SecureKey, error) {
	phrase := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, phrase); err != nil {
		return nil, nil, err
	}

	k := &SecureKey{
		phrase: phrase,
	}
	s, err := k.New(u)
	return s, k, err
}

// New prepares a secure string based on an existing key.
func (k *SecureKey) New(s string) (*SecureString, error) {
	cip, err := aes.NewCipher(k.phrase)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ss := &SecureString{
		nonce:  nonce,
		sealed: gcm.Seal(nil, nonce, []byte(s), nil),
	}

	return ss, nil
}

// Destroy purges a secure string.
func (s *SecureString) Destroy() {
	io.ReadFull(rand.Reader, s.sealed)
}

// String returns the readable string.
func (s *SecureString) String(k *SecureKey) (string, error) {
	cip, err := aes.NewCipher(k.phrase)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return "", err
	}

	bytes, err := gcm.Open(nil, s.nonce, s.sealed, nil)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
