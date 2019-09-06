package main

// type Conn interface {
//         // Read reads data from the connection.
//         // Read can be made to time out and return an Error with Timeout() == true
//         // after a fixed time limit; see SetDeadline and SetReadDeadline.
//         Read(b []byte) (n int, err error)
//
//         // Write writes data to the connection.
//         // Write can be made to time out and return an Error with Timeout() == true
//         // after a fixed time limit; see SetDeadline and SetWriteDeadline.
//         Write(b []byte) (n int, err error)
//
//         // Close closes the connection.
//         // Any blocked Read or Write operations will be unblocked and return errors.
//         Close() error
//
//         // LocalAddr returns the local network address.
//         LocalAddr() Addr
//
//         // RemoteAddr returns the remote network address.
//         RemoteAddr() Addr
//
//         // SetDeadline sets the read and write deadlines associated
//         // with the connection. It is equivalent to calling both
//         // SetReadDeadline and SetWriteDeadline.
//         //
//         // A deadline is an absolute time after which I/O operations
//         // fail with a timeout (see type Error) instead of
//         // blocking. The deadline applies to all future and pending
//         // I/O, not just the immediately following call to Read or
//         // Write. After a deadline has been exceeded, the connection
//         // can be refreshed by setting a deadline in the future.
//         //
//         // An idle timeout can be implemented by repeatedly extending
//         // the deadline after successful Read or Write calls.
//         //
//         // A zero value for t means I/O operations will not time out.
//         SetDeadline(t time.Time) error
//
//         // SetReadDeadline sets the deadline for future Read calls
//         // and any currently-blocked Read call.
//         // A zero value for t means Read will not time out.
//         SetReadDeadline(t time.Time) error
//
//         // SetWriteDeadline sets the deadline for future Write calls
//         // and any currently-blocked Write call.
//         // Even if write times out, it may return n > 0, indicating that
//         // some of the data was successfully written.
//         // A zero value for t means Write will not time out.
//         SetWriteDeadline(t time.Time) error
// }

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", ":2001")
	l, err := net.ListenTCP("tcp", addr)
	if nil != err {
		log.Fatalf("Listen fail: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if nil != err {
			log.Fatalf("Accept fail: %v", err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var iput string
	buf := make([]byte, 5)
	for {
		conn.Read(buf)
		input = string + string(buf)

		// 如果一次消息接收完成
		if bytes.Contains(buf, []byte("EOF")) {
			go biz(input, conn)
			input = ""
		}

		buf = make([]byte, 5)
	}
}

func biz(input string, conn net.Conn) {
	t, _ := time.Now().MarshalText()
	output := make([]byte, 0)

	conn.Write(output)
}
