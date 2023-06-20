package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num)

	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&num)
	newValue := pointer.Elem()

	fmt.Println("type of newValue:", newValue.Type())
	fmt.Println("settability of newValue:", newValue.CanSet())
	fmt.Println("settability of pointer:", pointer.CanSet())

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num)

	// 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
	// pointer = reflect.ValueOf(num)
	// newValue = pointer.Elem()
}
