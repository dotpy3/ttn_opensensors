package main

import (
	"bytes"
	"io"
)

// ReaderToString gets the content of a reader and returns it as a string
func ReaderToString(read io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(read)
	return buf.String()
}
