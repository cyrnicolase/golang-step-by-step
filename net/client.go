package main

import (
	// "bufio"
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.3.224:2000")
	if nil != err {
		fmt.Errorf("%v", err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	b, _ := ioutil.ReadAll(conn)

	fmt.Println(b)
}
