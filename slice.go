package main

import (
	"log"
)

func main() {
	s := make([]int, 0, 0)
	//s := []int{}
	s = append(s, 1)
	print(s)
	s = append(s, 1)
	print(s)

	s1 := make([]int, 10)
	s1[0] = 1
	print(s1)
	change(s1)
	print(s1)

	s2 := make([]int, 3)
	print(s2)
}

func change(s []int) {
	s = append(s, 2)
	print(s)
}

func print(s []int) {
	log.Printf("addr:%p, slice:%+v, len:%+v, cap:%+v", &s[0], s, len(s), cap(s))
}
