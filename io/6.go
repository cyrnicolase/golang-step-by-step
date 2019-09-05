package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	r3 := strings.NewReader("third reader\n")

	r4 := io.MultiReader(r1, r2, r3)
	if _, err := io.Copy(os.Stdout, r4); nil != err {
		log.Fatal(err)
	}
}
