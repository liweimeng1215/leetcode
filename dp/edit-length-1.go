// 求最小编辑距离递归解法 参考：https://mp.weixin.qq.com/s/uWzSvWWI-bWAV3UANBtyOw
package main

import (
	"fmt"
	"strings"
)

var visited = map[[2]int]int{}

func Min(a, b, c int) int {
	d := a
	if b < d {
		d = b
	}
	if c < d {
		d = c
	}
	return d
}

func dp(s1, s2 string, i, j int) int {
	ij := [2]int{i, j}
	var ans int
	if ans, ok := visited[ij]; ok {
		return ans
	}

	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if s1[i] == s2[j] {
		ans = dp(s1, s2, i-1, j-1)
	} else {
		ans = Min(dp(s1, s2, i, j-1)+1, // 增加
			dp(s1, s2, i-1, j)+1,   //删除
			dp(s1, s2, i-1, j-1)+1) // 替换
	}
	visited[ij] = ans
	return ans
}

func minEditLength(s1, s2 string) int {
	i, j := len(s1)-1, len(s2)-1
	return dp(s1, s2, i, j)
}

func main() {
	var input string
	fmt.Scanln(&input)
	d := strings.Split(input, ",")
	ans := minEditLength(d[0], d[1])
	fmt.Printf("%d\n", ans)
}
