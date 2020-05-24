package main

import "fmt"

func main() {
	fmt.Println(climbStairsV2(3))
}

//f(n) = f(n-1) + f(n-2)
//f(1) = 1, f(0) = 1
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}

func climbStairsV2(n int) int {
	mem := map[int]int{}
	return climb(n, mem)
}

func climb(n int, mem map[int]int) int {
	if _, ok := mem[n]; !ok {
		if n <= 1 {
			mem[n] = 1
		} else {
			mem[n] = climb(n-2, mem) + climb(n-1, mem)
		}
	}
	return mem[n]
}
