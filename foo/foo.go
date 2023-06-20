package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println(x())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(c())
}

func x() int {
	i := 0
	defer func() {
		i = 2
	}()
	return i
}

func a() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func b() (err error) {
	err = io.EOF
	defer func() {
		fmt.Println(err)
		err = io.EOF
	}()
	return nil
}

func c() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 9
}
