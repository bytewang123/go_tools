package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePathsWithObstacles(3, 2))
}

//d(m,n) = d(m-1,n)+d(m,n-1)
//d(0,n-1) = 1, d(m-1,0) = 1

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	cant := false
	for i := 0; i < m; i++ {
		if cant {
			path[i][0] = 0
			continue
		}
		if obstacleGrid[i][0] == 1 {
			path[i][0] = 0
			cant = true
		} else {
			path[i][0] = 1
		}
	}
	cant = false
	for j := 0; j < n; j++ {
		if cant {
			path[0][j] = 0
			continue
		}
		if obstacleGrid[0][j] == 1 {
			path[0][j] = 0
			cant = true
		} else {
			path[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				path[i][j] = 0
			} else {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}
	fmt.Println(path)
	return path[m-1][n-1]
}
