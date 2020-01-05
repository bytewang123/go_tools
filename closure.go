package main

import "fmt"

func main() {
	f := add()
	fmt.Printf("f(1) = %+v\n", f(1)) //1
	fmt.Printf("f(1) = %+v\n", f(1)) //1

	p := pp(0)
	fmt.Printf("p() = %+v\n", p()) //1
	fmt.Printf("p() = %+v\n", p()) //2
	fmt.Printf("p() = %+v\n", p()) //3
	fmt.Printf("p() = %+v\n", p()) //4
	fmt.Printf("p() = %+v\n", p()) //5

	lv := localVar()
	fmt.Printf("lv(0) = %+v\n", lv(0)) //0
	fmt.Printf("lv(1) = %+v\n", lv(1)) //1
	fmt.Printf("lv(2) = %+v\n", lv(2)) //3
	fmt.Printf("lv(3) = %+v\n", lv(3)) //6
	fmt.Printf("lv(4) = %+v\n", lv(4)) //10
}

func add() func(i int) int {
	num := 0
	return func(i int) int {
		return num + i
	}
}

//每次调用后，i的值会被保存
func pp(i int) func() int {
	return func() int {
		i++
		return i
	}
}

//每次调用后，局部变量sum会被保存
func localVar() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
