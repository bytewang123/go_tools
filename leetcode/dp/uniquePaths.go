package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 2))
}

//d(m,n) = d(m-1,n)+d(m,n-1)
//d(0,n-1) = 1, d(m-1,0) = 1
func uniquePaths(m int, n int) int {
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}
	for j := 0; j < n; j++ {
		path[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	fmt.Println(path)
	return path[m-1][n-1]
}
