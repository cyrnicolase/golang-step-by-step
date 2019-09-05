package main

import (
	"fmt"
	// "io"
	"log"
	"os"
)

func main() {

	b := make([]byte, 5)
	r := os.Stdin
	if _, err := r.Read(b); nil != err {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, string(b))
}
