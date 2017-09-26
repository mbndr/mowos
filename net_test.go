package mowos

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/mbndr/logo"
)

// TestReadSendBytes tests the ReadBytes and SendBytes function
func TestReadSendBytes(t *testing.T) {
	InitLog()
	Log.SetLevel(logo.DEBUG)
	msg := []byte("This is a test message")
	var receiver bytes.Buffer

	// write to buffer and read from buffer
	SendBytes(&receiver, msg)
	read, _ := ReadBytes(bufio.NewReader(&receiver))

	if string(read) != string(msg) {
		t.Error("bytes sent wrong")
	}
}

// TestReadSendBytesWithEncryption tests the ReadBytes and SendBytes function with encryption
func TestReadSendBytesWithEncryption(t *testing.T) {
	UsedCryptor = NewPSKCryptor(
		[]byte("5g9hwlkq4j9w8hpmbcc319d5ßdkde3d"),
		[]byte("PSKTEST"),
	)
	TestReadSendBytes(t)
}
