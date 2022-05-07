//LC88. 合并两个有序数组
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	debug("nums1: %v\n", nums1)
	debug("nums2: %v\n", nums2)
	var i, j, k = m - 1, n - 1, len(nums1) - 1
	for j >= 0 {
		if i < 0 {
			nums1[k] = nums2[j]
			j--
			k--
			continue
		}
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}
	debug("nums1: %v\n", nums1)
}
