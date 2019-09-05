package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	ReadDirectory(".")
}

// ReadDirectory 读目录
func ReadDirectory(dir string) {
	files, err := ioutil.ReadDir(dir)
	if nil != err {
		log.Fatal(err)
	}

	for _, file := range files {
		name := dir + "/" + file.Name()
		if file.IsDir() {
			ReadDirectory(name)
			continue
		}

		fmt.Println(name)
	}
}
