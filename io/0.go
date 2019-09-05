package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); nil != err {
		log.Fatal(err)
	}

	r1 := strings.NewReader("use bytes.Buffer")
	buf := new(bytes.Buffer)
	io.Copy(buf, r1)

	fmt.Println(buf.String())
}
