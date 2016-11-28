package main

import (
	"bytes"
	"io"
	"net/http"
)

// ReaderToString gets the content of a reader and returns it as a string
func ReaderToString(read io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(read)
	return buf.String()
}

// ErrorResponse indicates if a response is a failed request response
func ErrorResponse(rep *http.Response) bool {
	return (rep != nil && rep.StatusCode != 200 && rep.StatusCode != 201)
}
