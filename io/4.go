package main

import (
	// "bytes"
	"io"
	"os"
)

func main() {
	buf := make([]byte, 5)

	io.ReadFull(os.Stdin, buf)
	io.WriteString(os.Stdout, string(buf))

	io.WriteString(os.Stdout, "Hello world\n")
}
