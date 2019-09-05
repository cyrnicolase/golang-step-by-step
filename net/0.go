package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv4(114, 8, 8, 8))

	ips, err := net.LookupIP("www.baidu.com")
	if nil != err {
		fmt.Errorf("查找ip错误:%v", err)
	}
	fmt.Println(ips)

	ip := net.ParseIP("192.168.3.224")
	fmt.Println(ip.DefaultMask()) // 255.255.255.0

	fmt.Println(ip.IsInterfaceLocalMulticast())
	fmt.Println(ip.IsMulticast())
	fmt.Println(ip.IsLinkLocalUnicast())
	fmt.Println(ip.IsUnspecified())

	ip = net.ParseIP("0.0.0.0")
	fmt.Println(ip.IsUnspecified())

	ip = net.ParseIP("121.15.41.168")
	fmt.Println(ip.To16())
	fmt.Println(ip.To4())

	fmt.Println(net.IPv4Mask(255, 255, 255, 0))
}
