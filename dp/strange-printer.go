// LC664. 奇怪的打印机
// dp解法，dp[i][j]表示字符串s[i, j]打印的最少次数
// 当s[i] == s[j] 时，dp[i][j] = dp[i][j-1]
// 当s[i] != s[j] 时，dp[i][j] = min{ dp[i][k]+dp[k+1][j], k ∈ [i, j) }
package main

import "fmt"

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				minCount := n
				for k := i; k < j; k++ {
					if dp[i][k]+dp[k+1][j] < minCount {
						minCount = dp[i][k] + dp[k+1][j]
					}
				}
				dp[i][j] = minCount
			}
			debug("s[%v, %v]=%v, dp[i][j]=%v\n", i, j, s[i:j+1], dp[i][j])
		}
	}
	return dp[0][n-1]
}
