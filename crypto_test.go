package mowos

import (
	"reflect"
	"testing"
)

// TestPSKCryptor checks if the PSKCryptor is working
func TestPSKCryptor(t *testing.T) {
	c := &PSKCryptor{
		key: []byte("5g9hwlkq4j9w8hpmbcc319d5ßdkde3d"),
		// identity is only needed if sending
	}

	str := []byte("Only a string for testing encryption")

	encr, err := c.encrypt(str)
	if err != nil {
		t.Error(err)
	}

	decr, err := c.decrypt(encr)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(str, decr) {
		t.Error("Encrypted string isn't the same as the original one")
	}
}
