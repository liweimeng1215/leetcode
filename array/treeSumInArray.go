// 给出一个数组 A ，长度在0到1000之间，找出其中最大的数 d = a + b + c，其中 a,b,c,d 均属于 A
// 结果返回 d 的坐标
package main

import (
	"fmt"
	"sort"
)

type BySort []int

func (bs BySort) Less(i, j int) bool { return bs[i] < bs[j] }
func (bs BySort) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs BySort) Len() int           { return len(bs) }

func threeSumInArray(a []int) []int {
	N := len(a)
	if N < 4 {
		return []int{}
	}
	sort.Sort(BySort(a)) // O(N) = Nlog(N)
	for i := N - 1; i >= 3; i-- {
		for j := i - 1; j >= 2; j-- {
			t := a[i] - a[j]
			lp, rp := 0, j-1
			for lp < rp { // O(N) = N
				if a[lp]+a[rp] == t {
					return []int{a[lp], a[rp], a[j], a[i]}
				} else if a[lp]+a[rp] > t {
					rp--
				} else if a[lp]+a[rp] < t {
					lp++
				}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{3, 1, 2, 6, 0, 0, 11}
	ans := threeSumInArray(a)
	fmt.Printf("ans: %v\n", ans)
}
