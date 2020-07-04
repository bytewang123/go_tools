package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumTotal([][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}))
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, len(triangle[i]))
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			path[i][0] = triangle[i][0]
		} else {
			path[i][0] += path[i-1][0] + triangle[i][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < len(path[i]); j++ {
			if j == len(path[i])-1 {
				path[i][j] = path[i-1][j-1] + triangle[i][j]
			} else {
				if path[i-1][j] < path[i-1][j-1] {
					path[i][j] = path[i-1][j] + triangle[i][j]
				} else {
					path[i][j] = path[i-1][j-1] + triangle[i][j]
				}
			}
		}
	}
	last := len(path[m-1])
	min := path[m-1][0]
	for i := 0; i < last-1; i++ {
		if min > path[m-1][i+1] {
			min = path[m-1][i+1]
		}
	}
	fmt.Println(path[m-1])
	return min
}
