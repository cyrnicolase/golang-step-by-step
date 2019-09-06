package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去玩吧"))

	// func ContainsAny(b []byte, chars string) bool
	// 返回chars中的utf8转换后的字符是否存在b中
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "lol!"))
}
