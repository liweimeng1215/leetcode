// 给定一个目标值 n 以及整数数组 A，其中 A 中元素 0<= A[i] <=9，
// 求由 A 中元素组成，不大于 n 的最大整数
/*
	n = 23121 A = [2, 4, 9]
	ans = 22999
*/
// 贪心解法
// 从左向右遍历目标值 n ，指针为 i，n长度为N
// 每次指向当前位整数 b，从A中选择不大于 b 的最大数
// 直到当前匹配的数 n[0:i) > ans[0:i)，后续 ans[i:N-1] 可全部选择 A 中的最大值
package main

import (
	"fmt"
	"sort"
	"strconv"
)

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

type Solution struct {
	n    int
	A    []int
	maxA int
	minA int
}
type BySorted []int

func (bs BySorted) Less(i, j int) bool { return bs[i] < bs[j] }
func (bs BySorted) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs BySorted) Len() int           { return len(bs) }

func (so Solution) getNotBigger(b int) int {
	for i := range so.A {
		if so.A[i] > b {
			return i - 1
		}
	}
	return len(so.A) - 1
}

func (so *Solution) MaxNum(n int, A []int) int { // n >=0 , A != [] && 0<=A[i]<=9
	debug("solution start-------\n")
	debug("n: %v, A: %v\n", n, A)
	defer debug("end-------\n")

	so.n = n
	sort.Sort(BySorted(A))
	so.A = A
	so.maxA = A[len(A)-1]
	so.minA = A[0]
	if n < 10 && n < A[0] {
		return -1
	}

	strN := strconv.FormatInt(int64(so.n), 10)
	ans := ""
	flag := false
	i := 0
	for i < len(strN) {
		if flag == false { // 从左到右匹配全部相等
			id := so.getNotBigger(int(strN[i] - '0'))
			for id == -1 { // 回溯
				debug("[%v] id == -1\n", i)
				if i == 0 {
					flag = true
					break
				}
				last := strN[i-1] - '0'
				ans = ans[0 : i-1]
				id = so.getNotBigger(int(last) - 1)
				i--
			}
			if id != -1 {
				ans += strconv.FormatInt(int64(so.A[id]), 10)
				flag = so.A[id] < int(strN[i]-'0')
			}
		} else {
			ans += strconv.FormatInt(int64(so.maxA), 10)
		}
		debug("[%v], [flag == %v] ans: %v\n", i, flag, ans)
		i++
	}
	intAns, err := strconv.ParseInt(ans, 10, 0)
	if err != nil {
		debug("parse ans error, err info: %v\n", err)
		return -1
	}
	return int(intAns)
}

func main() {
	so := Solution{}

	n := 23121
	A := []int{2, 4, 9}
	ans := so.MaxNum(n, A)
	fmt.Printf("expected: 22999, ans: %v\n", ans)

	n1 := 22099
	A1 := []int{1, 2, 4, 9}
	ans1 := so.MaxNum(n1, A1)
	fmt.Printf("expected: 21999, ans: %v\n", ans1)

	n2 := 22099
	A2 := []int{2, 4, 9}
	ans2 := so.MaxNum(n2, A2)
	fmt.Printf("expected: 9999, ans: %v\n", ans2)

	n3 := 22
	A3 := []int{3, 4, 5}
	ans3 := so.MaxNum(n3, A3)
	fmt.Printf("expected: 5, ans: %v\n", ans3)

	n4 := 5
	A4 := []int{6, 7, 8}
	ans4 := so.MaxNum(n4, A4)
	fmt.Printf("expected: -1, ans: %v\n", ans4)

	n5 := 10
	A5 := []int{0}
	ans5 := so.MaxNum(n5, A5)
	fmt.Printf("expected: 0, ans: %v\n", ans5)

}
