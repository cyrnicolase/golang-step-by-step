package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// JoinHostPort
	addr := net.JoinHostPort("192.168.3.224", "2001")
	fmt.Println(addr)

	// LookupAddr
	// 通过给定的IP地址，反向查找映射到该IP的域名
	names, err := net.LookupAddr("127.0.0.1")
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(names)

	host, port, _ := net.SplitHostPort(addr)
	fmt.Println(host, port)
}
