package client

import (
	"bytes"
	"io"
	"os"
)

// ReadZip reads zip archive from file into memory.
func ReadZip(path string) (r io.Reader, err error) {
	d, err := os.ReadFile(path)
	if err != nil {
		return
	}
	r = bytes.NewReader(d)
	return
}
