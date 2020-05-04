package main

import (
	"fmt"
)

func main() {
	var err error
	r, err := addError()
	fmt.Println(r, err)
}

func addError() (int, error) {
	return 1, fmt.Errorf("test error")
}
