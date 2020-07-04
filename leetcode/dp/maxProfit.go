package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfitV2([]int{7, 1, 5, 3, 6, 4}))
	//fmt.Println(maxProfitV2([]int{7, 6, 4, 3, 9}))
}

//前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
func maxProfit(prices []int) int {
	profit := make([]int, len(prices))

	profit[0] = 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if prices[i]-min > profit[i-1] {
			profit[i] = prices[i] - min
		} else {
			profit[i] = profit[i-1]
		}
	}
	max := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			max = profit[0]
			continue
		}
		if max < profit[i] {
			max = profit[i]
		}
	}
	return max
}

//dp[n] = max(dp[n-1],dp[n-1]+arr[n]-min)
//dp[0] = 0
func maxProfitV2(prices []int) int {
	profit := make([]int, len(prices))
	profit[0] = 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if profit[i-1] > profit[i-1]+prices[i]-min {
			profit[i] = profit[i-1]
		} else {
			profit[i] = profit[i-1] + prices[i] - min
		}
	}

	max := 0
	fmt.Println(profit)
	for i := 0; i < len(profit); i++ {
		if max < profit[i] {
			max = profit[i]
		}
	}
	return max
}
