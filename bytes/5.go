package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 将切片中的所有空字符都去掉
	// 将紧邻的字符拼接在一起作为一个切片
	// 将多个子切片合并成一个大切片
	// 返回的是二维数组
	// func Fields(s []byte) [][]byte
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("   foo bar    ba z   ")))

	// [
	//    ['f', 'o', 'o'],
	//    ['b', 'a', 'r'],
	//    ['b', 'a'],
	//    ['z'],
	// ]
}
