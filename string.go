package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "b1111"
	if a[0] == 'b' {
		fmt.Println("equal")
	}
	fmt.Println("'b' type", reflect.TypeOf('b'))
	fmt.Println("string('b') type", reflect.TypeOf(string('b')))
	fmt.Println("string a[0] type", reflect.TypeOf(a[0]))
	fmt.Println(a[0] == 'b')
	fmt.Printf("%d,%d\n", a[0], 'b')
	for _, v := range a {
		fmt.Println("element type", reflect.TypeOf(v))
	}

	c := a
	fmt.Println("c", c)
	c[0] = c[1]
	fmt.Println(a)
	for k, v := range a {
		fmt.Printf("%+v %+v", k, v)
	}
}
