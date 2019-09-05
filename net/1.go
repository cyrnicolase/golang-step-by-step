package main

import (
	"bytes"
	"fmt"
	// "io"
	"log"
	"net"
)

func main() {
	// l net.Listener 接口
	l, err := net.Listen("tcp", ":2000")
	if nil != err {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Fatal(err)
		}

		// Handle the connection in a new goroutine
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently
		go func(c net.Conn) {

			// Echo all incoming data
			fmt.Println(c.RemoteAddr())
			buf := new(bytes.Buffer)
			buf.ReadFrom(c)
			fmt.Println(buf.String())

			c.Close()
		}(conn)
	}
}
