package mowos

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// currently used cryptor implementation
var UsedCryptor *PSKCryptor

// Cryptor can encrypt and decrypt data.
type Cryptor interface {
	// encrypts given data and appends metadata
	// if this is necessary for decryption (e.g. identity)
	encrypt(data []byte) ([]byte, error)
	// strips eventual metadata (e.g. identity)
	// and decrypts package
	decrypt(data []byte) ([]byte, error)
}

// PSKCryptor encrypts and decrypts data with a pre-shared-key.
// Attention: This doesn't follow any official RFCs.
type PSKCryptor struct {
	key      []byte
	identity []byte
}

// NewPSKCryptor returns a PSKCryptor
func NewPSKCryptor(key, identity []byte) *PSKCryptor {
	return &PSKCryptor{
		key:      key,
		identity: identity,
	}
}

// encrypt data
func (c *PSKCryptor) encrypt(plaintext []byte) ([]byte, error) {

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	// array with iv and plaintext length
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// slice of iv
	iv := ciphertext[:aes.BlockSize]

	// fill iv
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// encrypt bytes from plain to cipher
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// add identity
	ciphertext = append(c.identity, ciphertext...)

	return ciphertext, nil
}

// decrypt data
func (c *PSKCryptor) decrypt(ciphertext []byte) ([]byte, error) {

	// remove identity
	ciphertext = bytes.TrimPrefix(ciphertext, c.identity)

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("Data to short")
	}

	// get iv
	iv := ciphertext[:aes.BlockSize]

	// decrypt
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext[aes.BlockSize:], nil
}
