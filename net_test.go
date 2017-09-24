package mowos

import (
	"bufio"
	"bytes"
	"testing"
)

// TestReadSend tests the ReadBytes and SendBytes function
func TestReadBytes(t *testing.T) {
	msg := []byte("This is a test message")
	var receiver bytes.Buffer

	// write to buffer and read from buffer
	SendBytes(&receiver, msg)
	read, _ := ReadBytes(bufio.NewReader(&receiver))

	if string(read) != string(msg) {
		t.Error("bytes sent wrong")
	}
}
