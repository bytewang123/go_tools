package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPathSum([][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if i == 0 {
			path[i][0] = grid[i][0]
		} else {
			path[i][0] += path[i-1][0] + grid[i][0]
		}
	}
	for j := 0; j < n; j++ {
		if j == 0 {
			path[0][j] = grid[0][j]
		} else {
			path[0][j] += path[0][j-1] + grid[0][j]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if path[i-1][j] < path[i][j-1] {
				path[i][j] = path[i-1][j] + grid[i][j]
			} else {
				path[i][j] = path[i][j-1] + grid[i][j]
			}
		}
	}
	fmt.Println(path)
	return path[m-1][n-1]
}
