package main

import (
	"fmt"
)

type dog struct {
	name  string
	color string
}

func main() {
	d1 := dog{"tom", "red"}
	d2 := dog{"tom", "red"}
	if d1 == d2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
		fmt.Printf("%+v\t%+v\n", d1, d2)
	}
	ds := map[dog]string{}
	ds[d1] = "d1"
	ds[d2] = "d2"
	fmt.Println(ds)
}
