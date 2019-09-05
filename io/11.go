package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	buf, err := ioutil.ReadFile("./io/0.go")
	if nil != err {
		log.Fatal(err)
	}

	fmt.Printf("%s", buf)
}
