// LC42. 接雨水
// 从左向右扫描，应用单调递减栈存储墙的高度，遇到比栈顶元素更高的墙，弹出栈顶元素
// 计算雨水的承载量，循环到合适的位置将当前位置的墙存入单调栈中

package main

import (
	"container/list"
	"fmt"
)

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Wall struct {
	H  int
	Id int
}

func peek(l *list.List) int {
	return l.Front().Value.(int)
}

func trap(height []int) int {
	N := len(height)
	if N <= 2 {
		return 0
	}
	var l list.List
	ans := 0
	cur := 0
	for cur < N {
		for l.Front() != nil && height[cur] > height[peek(&l)] {
			p1 := peek(&l)
			l.Remove(l.Front())
			if l.Front() == nil {
				break
			}

			p2 := peek(&l)
			w := cur - p2 - 1
			h := min(height[p2], height[cur]) - height[p1]
			ans += w * h
			debug("[cur: %v h: %v], [p1: %v h: %v], [p2: %v h: %v]\n", cur, height[cur], p1, height[p1], p2, height[p2])
		}
		l.PushFront(cur)
		cur++
	}

	return ans
}
