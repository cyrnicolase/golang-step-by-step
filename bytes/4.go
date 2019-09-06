package main

import (
	"bytes"
	"fmt"
)

func main() {

	// func EqualFold(s, t []byte) bool 忽略大小写的情况比对
	fmt.Println(bytes.EqualFold([]byte("GO"), []byte("Go")))
}
