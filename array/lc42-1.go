// LC42. 接雨水
// 分层遍历，遇到墙开始接雨水，再遇到后续墙后计算接雨水的量，将其加到结果中
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
	N := len(height)
	maxHeight := 0
	for i := range height {
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}

	ans := 0
	for h := 1; h <= maxHeight; h++ {
		start := false
		temp := 0
		debug("[%v] water: ", h)
		for i := 0; i < N; i++ {
			if start == false && height[i] >= h { // 扫描到第一堵墙
				start = true
			}

			if start == true && height[i] >= h { // 扫描到后续的墙
				ans += temp
				temp = 0
			}

			if start == true && height[i] < h {
				temp++
				debug("[%v] ", i)
			}
		}
		debug("\n")
	}

	return ans
}
