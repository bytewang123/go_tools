package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//dp[i] = max(dp[i-1]+nums[i],nums[i])
func maxSubArray(nums []int) int {
	maxSub := make([]int, len(nums))
	maxSub[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxSub[i-1]+nums[i] {
			maxSub[i] = nums[i]
		} else {
			maxSub[i] = maxSub[i-1] + nums[i]
		}
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			max = maxSub[0]
		} else {
			if max < maxSub[i] {
				max = maxSub[i]
			}
		}
	}
	return max
}
