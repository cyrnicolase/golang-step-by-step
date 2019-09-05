package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	printall := func(r io.Reader) {
		b, err := ioutil.ReadAll(r)
		if nil != err {
			log.Fatal(err)
		}

		fmt.Printf("%s", b)
	}

	printall(&buf)
	printall(tee)
}
