// LC33. 搜索旋转排序数组
// 将数组一分为二，一半为有序，一半为无序
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d == true {
		fmt.Printf(format, a...)
	}
}

func searchInSorted(nums []int, target, lo, hi int) int {
	for lo <= hi {
		mi := (lo + hi) / 2
		debug("[search in sorted], lo: %v, hi: %v, mi: %v, nums[mi]: %v, t: %v\n", lo, hi, mi, nums[mi], target)
		if nums[mi] == target {
			return mi
		} else if nums[mi] < target {
			lo = mi + 1
		} else if nums[mi] > target {
			hi = mi - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	debug("[search] nums: %v, t: %v\n", nums, target)
	var lo, hi = 0, len(nums) - 1
	for lo <= hi {
		mi := (lo + hi) / 2
		debug("[search], lo: %v, hi: %v, mi: %v, nums[mi]: %v, t: %v\n", lo, hi, mi, nums[mi], target)
		if nums[mi] == target {
			return mi
		} else {
			// left sorted and target in
			if lo <= mi-1 {
				if nums[lo] <= nums[mi-1] && nums[lo] <= target && target <= nums[mi-1] {
					return searchInSorted(nums, target, lo, mi-1)
				} else if nums[lo] <= nums[mi-1] && !(nums[lo] <= target && target <= nums[mi-1]) {
					lo = mi + 1
				}
			}
			if mi+1 <= hi {
				// right sorted and target in
				if nums[mi+1] <= nums[hi] && nums[mi+1] <= target && target <= nums[hi] {
					return searchInSorted(nums, target, mi+1, hi)
				} else if nums[mi+1] <= nums[hi] && !(nums[mi+1] <= target && target <= nums[hi]) {
					hi = mi - 1
				}
			}
			if lo == hi {
				break
			}
		}
	}
	return -1
}
