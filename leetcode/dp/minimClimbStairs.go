package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	//fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}

//d(n) = min(d(n-1),d(n-2)+a[n])

func minCostClimbingStairs(cost []int) int {
	mem := map[int]int{}
	return costClimb(len(cost)-1, cost, mem)
}

func costClimb(n int, cost []int, mem map[int]int) int {
	if _, ok := mem[n]; !ok {
		if n == 0 || n == 1 {
			mem[n] = cost[n]
		} else {
			costClimbLast2 := costClimb(n-2, cost, mem)
			costClimbLast1 := costClimb(n-1, cost, mem)
			if costClimbLast1 < costClimbLast2+cost[n] {
				mem[n] = costClimbLast1
			} else {
				mem[n] = costClimbLast2 + cost[n]
			}
		}
	}
	fmt.Println(mem)
	return mem[n]
}
