// 求最小编辑距离 自底向上解法
package main

import (
	"fmt"
	"strings"
)

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

func minEditLength(s1, s2 string) int {
	var dp = make([][]int, len(s2)+1)
	for i := range dp {
		dp[i] = make([]int, len(s1)+1)
	}

	for j := 1; j <= len(s1); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(s2); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(s2); i++ {
		for j := 1; j <= len(s1); j++ {
			if s2[i-1] == s1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i][j-1]+1, dp[i-1][j-1]+1, dp[i-1][j]+1)
			}
		}
	}
	return dp[len(s2)][len(s1)]
}

func main() {
	var input string
	fmt.Scanln(&input)
	d := strings.Split(input, ",")
	ans := minEditLength(d[0], d[1])
	fmt.Printf("%d\n", ans)
}
