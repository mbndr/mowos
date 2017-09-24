// THIS DOESNT WORK / ISNT USED
package mowos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Cryptor can encrypt and decrypt data
type Cryptor interface {
	encrypt(data []byte) ([]byte, error)
	decrypt(data []byte) ([]byte, error)
}

// PSKCryptor encrypts and decrypts data with a pre-shared-key.
// Attention: This doesn't follow any official RFCs.
type PSKCryptor struct {
	key []byte
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

	return ciphertext, nil
}

// decrypt data
func (c *PSKCryptor) decrypt(ciphertext []byte) ([]byte, error) {

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
