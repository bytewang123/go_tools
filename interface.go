package main

import (
	"fmt"
)

type Animal interface {
	Bark() string
	Walk() string
}

type Dog struct {
	name string
}

func (dog Dog) Bark() string {
	fmt.Println(dog.name + ":wan wan wan!")
	return "wan wan wan"
}

func (dog Dog) Walk() string {
	fmt.Println(dog.name + ":walk to park!")
	return "walk to park"
}

func main() {
	var animal Animal

	fmt.Println("animal value is:", animal)    //animal value is: <nil>
	fmt.Printf("animal type is: %T\n", animal) //animal type is: <nil>

	animal = Dog{"旺财"}
	animal.Bark() //旺财:wan wan wan!
	animal.Walk() //旺财:walk to park!

	fmt.Println("animal value is:", animal)    //animal value is: {旺财}
	fmt.Printf("animal type is: %T\n", animal) //animal type is: main.Dog
}
