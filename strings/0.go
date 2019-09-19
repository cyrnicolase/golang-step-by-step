package main

import (
	"fmt"
	"strings"
)

// END_OF_FILE 结束符
const END_OF_FILE = "EOF"

func main() {

	s1 := "   Hello \n\r  \n"
	s1 = strings.TrimSpace(s1)
	fmt.Printf("[%s]\n", s1)

	fmt.Printf("[%s]\n", strings.TrimRight("Hello world!EOF", END_OF_FILE))
}
