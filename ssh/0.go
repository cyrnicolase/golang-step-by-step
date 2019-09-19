package main

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	// var hostKey ssh.PublicKey

	config := &ssh.ClientConfig{
		User: "www",
		Auth: []ssh.AuthMethod{
			ssh.Password("123456"),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	client, err := ssh.Dial("tcp", "192.168.0.11:22", config)
	if nil != err {
		log.Fatal("fail to dial: ", err)
	}

	session, err := client.NewSession()
	if nil != err {
		log.Fatal("fail to newsession: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b

	if err := session.Run("/bin/ls -al /data/wwwroot"); nil != err {
		log.Fatal("fail run command: ", err)
	}

	fmt.Println(b.String())
}
