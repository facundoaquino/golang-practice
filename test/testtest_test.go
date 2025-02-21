package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	expected := "Hello, W0orld!\n"
	if got := captureOutput(main); got != expected {
		t.Errorf("main() = %v, want %v", got, expected)
	}
}

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	defer func() { os.Stdout = stdout }()
	os.Stdout = w

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
