package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 判断是否包含
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))

	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))

	fmt.Println(bytes.Contains([]byte(""), []byte("")))

	str := "Hello world"
	fmt.Println(bytes.Contains([]byte(str), []byte("ello"))) // 字节切片包含
	fmt.Println(strings.Contains(str, "ello"))               // 字符串包含
}
