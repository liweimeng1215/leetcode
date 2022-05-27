// LC1313. 解压缩编码列表
// 模拟解码
// 注：先申请解码长度的数组，可避免数组动态扩容的开销
package main

import "fmt"

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func decompressRLElist(nums []int) []int {
	n := len(nums)
	if n <= 0 {
		return []int{}
	}

	len := 0
	for i := 0; 2*i+1 < n; i++ {
		len += nums[2*i]
	}
	ans := make([]int, len)

	j := 0
	for i := 0; 2*i+1 < n; i++ {
		freq, val := nums[2*i], nums[2*i+1]
		for k := j; k < j+freq; k++ {
			ans[k] = val
		}
		debug("freq: %v, val: %v, ans: %v\n", freq, val, ans)
		j = j + freq
	}
	return ans
}
