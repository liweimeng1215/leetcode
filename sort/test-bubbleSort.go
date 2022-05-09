// 冒泡排序
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

// arr := [5, 4, 3, 6, 7]
func bubbleSort(arr []int) {
	n := len(arr)
	lastJ := 0
	borderJ := n - 1
	for i := 0; i < n; i++ {
		debug("[%v] arr: %v\n", i, arr)
		sorted := true                 //  tip 2: 默认有序
		for j := 0; j < borderJ; j++ { // tip 1: j 每次遍历的边界和最近一次交换的位置有关系
			if arr[j+1] < arr[j] {
				sorted = false
				arr[j+1], arr[j] = arr[j], arr[j+1]
				lastJ = j // tip 1: 记录最新的交换位置
			}
		}
		if sorted { // tip 2
			break
		} else {
			borderJ = lastJ
		}
	}
	fmt.Printf("arr: %v\n", arr)
}

func main() {
	bubbleSort([]int{7, 5, 4, 3, 2, 6})
}
