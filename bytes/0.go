package main

import (
	"bytes"
	"fmt"
)

func main() {
	var a, b []byte
	if 0 > bytes.Compare(a, b) {
		fmt.Println("a less b")
	}
	if 0 == bytes.Compare(a, b) {
		fmt.Println("a equal b")
	}
	if 1 == bytes.Compare(a, b) {
		fmt.Println("a greater b")
	}
}
