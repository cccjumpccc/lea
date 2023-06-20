package main

import (
	"fmt"
	"reflect"
)

func main() {
	ifa := new(Person)
	var ifb Person = Student{name: "halfrost"}
	// 未绑定具体变量的接口类型
	fmt.Println(reflect.TypeOf(ifa).Elem().Name())
	fmt.Println(reflect.TypeOf(ifa).Elem().Kind().String())
	// 绑定具体变量的接口类型
	fmt.Println(reflect.TypeOf(ifb).Name())
	fmt.Println(reflect.TypeOf(ifb).Kind().String())
}
