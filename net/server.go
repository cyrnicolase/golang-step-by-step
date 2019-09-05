package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	if nil != err {
		fmt.Errorf("%v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			fmt.Errorf("%v", err)
		}

		// handler
		go func(c net.Conn) {
			defer c.Close()

			b := make([]byte, 1024)
			c.Read(b)

			fmt.Println(string(b))

			io.Copy(c, c)
		}(conn)
	}
}
