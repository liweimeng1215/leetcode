// LC56. 合并区间
// 首先将intervals排序（根据 interval 的起始位置）
// 然后用 i 指针指向已加入到结果的最后一个元素
// j 从第二个元素遍历intervals，判断是否需要将当前元素合并到结果中
//
package main

import "sort"

type SortSlices [][]int

func (ss SortSlices) Len() int           { return len(ss) }
func (ss SortSlices) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }
func (ss SortSlices) Less(i, j int) bool { return ss[i][0] < ss[j][0] }

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}

	sort.Sort(SortSlices(intervals))

	ans := [][]int{intervals[0]}
	i := 0
	for j := 1; j < n; j++ {
		lastStart, lastEnd := ans[i][0], ans[i][1]
		curStart, curEnd := intervals[j][0], intervals[j][1]
		if curStart <= lastEnd { // merge
			temp := []int{lastStart, max(lastEnd, curEnd)}
			ans[i] = temp
		} else { // not merge
			temp := []int{curStart, curEnd}
			ans = append(ans, temp)
			i++
		}
	}
	return ans
}
