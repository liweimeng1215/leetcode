// LC53. 最大子数组和，动态规划优化版本
package main

func maxSubArray(nums []int) int {
	var old = nums[0]
	var ans = old

	for i := 1; i < len(nums); i++ {
		if old+nums[i] > nums[i] {
			old = old + nums[i]
		} else {
			old = nums[i]
		}
		if old > ans {
			ans = old
		}
	}
	return ans
}
