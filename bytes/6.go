package main

import (
	"bytes"
	"fmt"
)

func main() {

	// bytes的很多功能跟strings 都很相像
	// []byte 和 string之间也可以直接进行转换
	fmt.Println("HasPrefix")
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))

	fmt.Println("HasSuffix")
	fmt.Println(bytes.HasSuffix([]byte("Gopher"), []byte("")))
	fmt.Println(bytes.HasSuffix([]byte("Gopher"), []byte("r")))
	fmt.Println(bytes.HasSuffix([]byte("Gopher"), []byte("er")))
	fmt.Println(bytes.HasSuffix([]byte("Gopher"), []byte("e")))

	fmt.Println("Index")
	// 查找待查询字节切片首次出现在被查询字节切片的位置
	fmt.Println(bytes.Index([]byte("Gopher loves go"), []byte("G")))
	fmt.Println(bytes.Index([]byte("Gopher loves go"), []byte("o")))
	fmt.Println(bytes.Index([]byte("Gopher loves go"), []byte("he")))
	fmt.Println(bytes.Index([]byte("Gopher loves go"), []byte("hee")))

	fmt.Println("IndexByte")
	fmt.Println(bytes.IndexByte([]byte("Hello"), 'e'))
	fmt.Println(bytes.IndexByte([]byte("Hello"), 'x'))
}
