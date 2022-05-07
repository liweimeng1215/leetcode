// LC53. 最大子数组和，动态规划优化版本
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d == true {
		fmt.Printf(format, a...)
	}
}

func maxSubArray(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	debug("dp: %v", dp)

	ans := dp[0]
	for i := range dp {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
