package mowos

import (
	"bufio"
	"bytes"
	"io"
)

var stopChar = []byte("\r\n\r\n")

// ReadBytes reads the bytes from a connection until the end symbol appears
// and decrypts the data.
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

	data := bytes.TrimSuffix(buf.Bytes(), stopChar)

	Log.Debugf("RECEIVING (RAW): %#v", string(data))

	// decrypt if a cryptor is set
	if UsedCryptor != nil {
		decr, err := UsedCryptor.decrypt(data)
		if err != nil {
			return nil, err
		}
		data = decr
	}

	Log.Debugf("RECEIVING (DECR): %#v", string(data))

	return data, nil
}

// SendBytes sends data over tcp after encrypting
func SendBytes(w io.Writer, data []byte) error {

	// encrypt if a cryptor is set
	if UsedCryptor != nil {
		encr, err := UsedCryptor.encrypt(data)
		if err != nil {
			return err
		}
		data = encr
	}

	w.Write(data)
	w.Write(stopChar)

	return nil
}
