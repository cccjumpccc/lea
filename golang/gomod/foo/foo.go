package main

import (
	"log"

	"github.com/fatih/structs"
)

type Dog struct {
	Name string `json:"aaa"`
	Age  int    `json:"bbb"`
}

func main() {
	s := make(map[string]interface{})
	s["name"] = "wang"
	s["age"] = 10
	for k := range s {
		log.Println(k)
	}
	structs.Map(Dog{})
}
