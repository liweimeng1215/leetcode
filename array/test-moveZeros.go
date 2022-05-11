// LC283. 移动零
// 双指针解法
package main

import "fmt"

func moveZeroes(nums []int) {
	var i, j = 0, 0
	for j < len(nums) { // [i, j)之间为 0 元素
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
}

func main() {
	nums := []int{0, 1, 2, 0, 4, 0, 0, 7}
	moveZeroes(nums)
	fmt.Printf("nums: %v", nums)
}
