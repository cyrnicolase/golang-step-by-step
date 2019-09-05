package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r1 := strings.NewReader("first line\n")  // 创建Reader
	r2 := strings.NewReader("second line\n") // 创建Reader

	buf := make([]byte, 8) // 声明缓冲区
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); nil != err {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); nil != err {
		log.Fatal(err)
	}

	fmt.Println(string(buf))
}
