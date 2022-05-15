// LC11. 盛最多水的容器
// 双指针解法，i，j 分别指向数组的左右两端
// 向中心移动短板：1）移动长版会减小区域的面积；2）移动短板，可能会增大区域的面积
//
package main

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1

	ans := 0
	for i < j {
		temp := min(height[i], height[j]) * (j - i)
		if temp > ans {
			ans = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans
}
