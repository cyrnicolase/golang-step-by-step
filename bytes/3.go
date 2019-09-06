package main

import (
	"bytes"
	"fmt"
)

func main() {
	// func Count(s, sep []byte) int
	// 返回切片sep 在s中出现的次数
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))

	// 切片出现2次
	fmt.Println(bytes.Count([]byte("hello hero"), []byte("he")))

	// 待查询切片如果是空，那么返回原切片长度+1
	fmt.Println(bytes.Count([]byte("cheese"), []byte("")))

	// 使用的是UTF8的长度，而非字节长度
	fmt.Println(bytes.Count([]byte("我是中国人,我爱吃川菜!"), []byte("")))
	fmt.Println(bytes.Count([]byte("我是中国人,我爱吃川菜!"), []byte("我")))
}
