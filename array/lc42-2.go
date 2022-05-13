// LC42. 接雨水
// 按行计算能够承载的雨水量
// 找到当前行左侧以及右侧最高的墙，当前能承载的雨水根据这两者与当前墙的关系决定
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
	ans := 0
	for i := 1; i <= N-2; i++ { // 找到 1~N-2 位置上 左最高墙 以及 右最高墙，计算能承载的雨水量
		j, k := i-1, i+1
		leftMax, rightMax := 0, 0
		debug("[%v] ", i)
		for 0 <= j {
			if height[j] > leftMax {
				leftMax = height[j]
			}
			j--
		}
		for k <= N-1 {
			if height[k] > rightMax {
				rightMax = height[k]
			}
			k++
		}
		debug("[l: %v m: %v r: %v]", leftMax, height[i], rightMax)
		if leftMax > height[i] && rightMax > height[i] {
			if leftMax < rightMax {
				ans += leftMax - height[i]
				debug("[%v]", leftMax-height[i])
			} else {
				ans += rightMax - height[i]
				debug("[%v]", rightMax-height[i])
			}
		}
		debug("\n")
	}
	return ans
}
