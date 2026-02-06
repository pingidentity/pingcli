// Copyright © 2026 Ping Identity Corporation

package testutils

import (
	"context"
	"io"
	"os"
	"time"
)

// CaptureStdout executes a function and returns its standard output as a string.
// If the function takes longer than 30 seconds to complete, it returns an empty string.
func CaptureStdout(f func()) string {
	originalStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return ""
	}
	defer func() { _ = r.Close() }()

	defer func() { os.Stdout = originalStdout }()
	os.Stdout = w

	outC := make(chan string, 1)
	go func() {
		b, _ := io.ReadAll(r)
		outC <- string(b)
	}()

	done := make(chan struct{})
	go func() {
		f()
		_ = w.Close()
		close(done)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		_ = w.Close()

		return ""
	case <-done:
		return <-outC
	}
}
