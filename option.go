package main

import "fmt"

//用于初始化参数不定的情况，可以传入任意个参数
//可以给每个参数一个初始默认值，
//如果传入了ServerOption函数，就可以自定义这个字段的值，没传入这个字段的ServerOption则使用默认值

type options struct {
	a int
	b string
}

//定义了ServerOption为一个函数，参数是*options
//这样对参数*options赋值，就会一直被保存
type ServerOption func(*options)

func NewOption(opt ...ServerOption) *options {
	r := new(options)
	for _, o := range opt {
		o(r)
	}
	return r
}

func InitA(a int) ServerOption {
	return func(o *options) {
		o.a = a
	}
}

func InitB(b string) ServerOption {
	return func(o *options) {
		o.b = b
	}
}

func main() {
	initFuncs := []ServerOption{
		InitA(100),
		InitB("hello"),
	}
	opts := NewOption(initFuncs...)
	fmt.Printf("opts = %+v\n", opts)
}
