package main

import (
	"fmt"
	"time"
)

func main() {
	today()
}

func today() {
	t := time.Now().Format("2006-01-02")
	fmt.Println(t)
}
