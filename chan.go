package main

import (
	"fmt"
)

func main() {
	case4()
}

func case1() {
	c := make(chan int)
	//只有发送没有接收，死锁
	c <- 1
}

func case2() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

func case3() {
	c := make(chan int)
	go func() {
		//有时可以取到，有时无法取到
		fmt.Println(<-c)
	}()
	c <- 1
}

func case4() {
	//只有接收没有发送,死锁
	c := make(chan int)
	<-c
}
