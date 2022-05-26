// LC139. 单词拆分
// dp解法
// dp[i] --> s[0, i) 可以被 wordDict 中的word拼接
// dp[i] = true		if s[j, i) ∈ wordDict AND dp[j] == true
//       = false	else
package main

import "fmt"

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	wordMap := map[string]struct{}{}
	for _, w := range wordDict {
		wordMap[w] = struct{}{}
	}

	for i := 1; i <= n; i++ {
		ans := false
		for _, w := range wordDict {
			j := i - len(w)
			if j < 0 {
				continue
			}
			temp := s[j:i]
			if _, ok := wordMap[temp]; ok && dp[j] {
				ans = true
				break
			}
		}
		dp[i] = ans
	}
	debug("dp: %v\n", dp)
	return dp[n]
}
