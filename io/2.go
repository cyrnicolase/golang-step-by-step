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
	r := strings.NewReader("some io.Reader stream to be used")

	if _, err := io.CopyN(os.Stdout, r, 5); nil != err {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, "\n")

	if _, err := io.CopyN(os.Stdout, r, 10); nil != err {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, "\n")

	buf := new(bytes.Buffer)
	io.CopyN(buf, r, 10)
	fmt.Fprintf(os.Stdout, "偏移字符串：%s", buf.String())
}
