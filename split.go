package main

import (
	"fmt"
	"strings"
)

func main() {
	input := `xxxxx:yyyyy:zzz.aaa.bbb.cc:dd:ee:ffMAX111`
	a := strings.FieldsFunc(input, Split)
	fmt.Println(a) // {xxxxx yyyyy zzz aaa bbb cc dd}
}
func Split(r rune) bool {
	return r == ':' || r == '.' || r == 'MAX'
}
