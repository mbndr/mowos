package mowos

import (
	"bufio"
	"bytes"
	"io"
	"net"
)

var stopChar = []byte("\r\n\r\n")

// ReadBytes reads the bytes from a connection until the end symbol appears
// Reader as param so that unit tests will be easier.
func ReadBytes(r *bufio.Reader) ([]byte, error) {

	var buf bytes.Buffer

	for {
		b, err := r.ReadByte()

		if err == io.EOF { // connection closed
			break
		} else if err == nil { // still reading
			buf.WriteByte(b)
			// end bytes reached
			if bytes.HasSuffix(buf.Bytes(), stopChar) {
				break
			}
		} else { // receiving failed
			return nil, err
		}

	}

	trimmed := bytes.TrimSuffix(buf.Bytes(), stopChar)
	return trimmed, nil
}

// SendBytes sends data over tcp
func SendBytes(conn net.Conn, data []byte) {
	conn.Write(data)
	conn.Write(stopChar)
}
