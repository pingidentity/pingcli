// Copyright © 2025 Ping Identity Corporation

package testutils

import (
	"io"
	"os"
)

// CaptureStdout executes a function and returns its standard output as a string.
func CaptureStdout(f func()) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()

	defer func() { os.Stdout = originalStdout }()

	os.Stdout = w

	outC := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		outC <- string(b)
	}()

	f()

	err := w.Close()
	if err != nil {
		return ""
	}

	return <-outC
}
