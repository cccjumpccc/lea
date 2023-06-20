package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("/dev/shm/hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
