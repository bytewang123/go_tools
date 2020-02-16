package main

import (
	"errors"
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

type TagTest struct {
	Name string `json:"name_json"`
	Age  int    `json:"age_json"`
}

type T struct{}

func (t *T) Add(a, b int) {
	fmt.Printf("a + b is %+v\n", a+b)
}

func (t *T) AddRetErr(a, b int) (int, error) {
	if a+b < 10 {
		return a + b, errors.New("total lt 10")
	}
	return a + b, nil
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

	t := TagTest{Name: "tom", Age: 10}
	rtt := reflect.TypeOf(t)
	//rtv := reflect.ValueOf(t)
	for i := 0; i < rtt.NumField(); i++ {
		field := rtt.Field(i)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("tag is %+v, value is %+v\n", json, field.Tag.Get("json"))
		}
	}

	funcName := "Add"
	typeT := &T{}
	a := reflect.ValueOf(5)
	b := reflect.ValueOf(6)
	in := []reflect.Value{a, b}
	reflect.ValueOf(typeT).MethodByName(funcName).Call(in)

	funcName = "AddRetErr"
	ret := reflect.ValueOf(typeT).MethodByName(funcName).Call(in)
	fmt.Printf("ret is %+v\n", ret)

	for i := 0; i < len(ret); i++ {
		fmt.Printf("ret index:%+v, type:%+v, value:%+v\n", i, ret[i].Kind(), ret[i].Interface())
	}

	if v, ok := ret[1].Interface().(error); ok {
		fmt.Printf("v is %+v\n", v)
	}

	t1 := TagTest{}
	tV := reflect.ValueOf(t)
	t1V := reflect.ValueOf(&t1)
	fmt.Printf("tV:%+v\n", tV)
	for i := 0; i < tV.NumField(); i++ {
		val := tV.Field(i)
		if t1V.Elem().CanSet() {
			t1V.Elem().Field(i).Set(val)
		}
	}
	fmt.Printf("t1 is %+v\n", t1)

	ta := 10
	vta := reflect.ValueOf(&ta)
	if vta.Elem().CanSet() {
		vta.Elem().Set(reflect.ValueOf(11))
	}
	fmt.Println("cant set")
	fmt.Printf("ta is :%+v\n", ta)

	ts := []int{1, 2, 3}
	tsV := reflect.ValueOf(ts)
	if tsV.Index(0).CanSet() {
		tsV.Index(0).Set(reflect.ValueOf(10))
	}
	fmt.Printf("ts is %+v\n", ts)

	tsA := [3]int{1, 2, 3}
	tsAv := reflect.ValueOf(&tsA)
	if tsAv.Elem().Index(0).CanSet() {
		tsAv.Elem().Index(0).Set(reflect.ValueOf(10))
	}
	fmt.Printf("tsA is %+v\n", tsA)
}
