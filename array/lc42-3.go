// LC42. 接雨水
// 基于 lc42-2 进行优化，不必每个点通过遍历去查找对应的 leftMax 以及 rightMax
// 通过扫描一遍 height， 得到 leftMax 和 rightMax 数组
// leftMax[i] = leftMax[i-1] if height[i] < leftMax[i-1]
//			= height[i] else
// rightMax[j] = rightMax[j+1] if height[j]< rightMax[j+1]
//			= height[i] else
//
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func trap(height []int) int {
	debug("height: %v\n", height)
	N := len(height)
	if N <= 2 {
		return 0
	}
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0], leftMax[1], rightMax[N-2], rightMax[N-1] = 0, height[0], height[N-1], 0
	for i := 2; i < N; i++ {
		j := N - 1 - i
		if leftMax[i-1] < height[i-1] {
			leftMax[i] = height[i-1]
		} else {
			leftMax[i] = leftMax[i-1]
		}

		if height[j+1] < rightMax[j+1] {
			rightMax[j] = rightMax[j+1]
		} else {
			rightMax[j] = height[j+1]
		}
	}
	debug("leftMax: %v\n", leftMax)
	debug("rightMax: %v\n", rightMax)
	ans := 0
	for i := 1; i <= N-2; i++ {
		if leftMax[i] > height[i] && height[i] < rightMax[i] {
			if leftMax[i] < rightMax[i] {
				ans += leftMax[i] - height[i]
			} else {
				ans += rightMax[i] - height[i]
			}
		}
	}
	return ans
}
