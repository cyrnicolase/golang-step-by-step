package main

import (
	"bytes"
	// "io"
	// "log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader to be used\n")
	r.WriteTo(os.Stdout)

	buf, err := r.ReadByte()
	if nil != err {
		log.Fatal(err)
	}

}
