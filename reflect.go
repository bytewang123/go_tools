package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name     string
	Sex      string
	Age      int
	PhoneNum string
	School   string
	City     string
}

func main() {
	p1 := Person{
		Name:     "tom",
		Sex:      "male",
		Age:      10,
		PhoneNum: "1000000",
		School:   "spb-kindergarden",
		City:     "cq",
	}

	rv := reflect.ValueOf(p1)
	rt := reflect.TypeOf(p1)
	if rv.Kind() == reflect.Struct {
		for i := 0; i < rt.NumField(); i++ {
			fmt.Printf("field:%+v,value:%+v\n", rt.Field(i).Name, rv.Field(i))
		}
	}
	if f, ok := rt.FieldByName("Age"); ok {
		fmt.Printf("field:%+v,value:%+v\n", f.Name, rv.FieldByName("Age"))
	}

	fmt.Printf("type:%+v, value:%+v\n", rt, rv)
	fmt.Printf("kind is %+v\n", rt.Kind())
	fmt.Printf("kind is %+v\n", rv.Kind())
}
