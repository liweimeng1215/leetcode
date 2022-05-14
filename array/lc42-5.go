// LC42. 接雨水
// 双指针解法 参考 https://leetcode-cn.com/problems/trapping-rain-water/solution/jie-yu-shui-by-leetcode/327718
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func trap(height []int) int {
	N := len(height)
	if N <= 2 {
		return 0
	}
	left, right := 0, N-1
	leftMax, rightMax := 0, 0
	ans := 0
	for left <= right {
		if leftMax < rightMax { // 可参考 lc32-3.go leftMax 和 rightMax 都是单调递增的
			if leftMax > height[left] {
				ans += leftMax - height[left]
			}
			leftMax = max(leftMax, height[left])
			left++
		} else {
			if rightMax > height[right] {
				ans += rightMax - height[right]
			}
			rightMax = max(rightMax, height[right])
			right--
		}
	}
	return ans
}
