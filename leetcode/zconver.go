package main

import "fmt"

//https://leetcode-cn.com/problems/zigzag-conversion/

func main() {
	convert("APPLE", 3)
}

func convert(s string, numRows int) string {
	min := 0
	if numRows == 1 {
		return s
	}
	if len(s) < numRows {
		return s
	} else {
		min = numRows
	}

	rows := [][]rune{}
	for i := 0; i < min; i++ {
		rows = append(rows, []rune(""))
	}
	runes := []rune(s)
	idx := 0
	up := false
	for _, c := range runes {
		rows[idx] = append(rows[idx], c)
		if idx == min-1 {
			up = true
		} else if idx == 0 {
			up = false
		}
		if up {
			idx--
		} else {
			idx++
		}
	}
	result := []rune{}
	for _, row := range rows {
		for _, c := range row {
			result = append(result, c)
		}
	}
	fmt.Printf("result:%+v\n", string(result))
	return string(result)
}
