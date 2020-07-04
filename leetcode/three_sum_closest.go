package main

import (
	"fmt"
)

func main() {
	result := threeSumClosest([]int{1, 1, 1, 0}, 100)
	fmt.Println(result)
}

func threeSumClosest(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
	idx := 0
	min := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			min = abs(target, nums[i])
			continue
		}
		if abs(nums[i], target) < min {
			min = abs(nums[i], target)
			idx = i
		}

	}
	fmt.Println(idx)
	if idx < 1 {
		return nums[0] + nums[1] + nums[2]
	} else if idx >= len(nums)-3 {
		return nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
	} else if idx == 1 {
		result := 0
		total1 := nums[idx-1] + nums[idx] + nums[idx+1]
		total2 := nums[idx] + nums[idx+1] + nums[idx+2]
		if abs(target, total2) > abs(target, total1) {
			result = total1
		} else {
			result = total2
		}
		return result
	} else {
		total1 := nums[idx-2] + nums[idx-1] + nums[idx]
		total2 := nums[idx-1] + nums[idx] + nums[idx+1]
		total3 := nums[idx] + nums[idx+1] + nums[idx+2]
		result := 0
		if abs(target, total2) > abs(target, total1) {
			result = total1
		} else {
			result = total2
		}
		if abs(target, total3) < abs(target, result) {
			result = total3
		}
		return result
	}
}

func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	} else {
		return a - b
	}
}
