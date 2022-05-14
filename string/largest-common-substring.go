// 字符串s与其逆字符串的最长公共子串
package main

func reverse(s string) string {
	N := len(s)
	ans := make([]byte, N)
	for i := 0; i < N; i++ {
		j := N - 1 - i
		ans[j] = s[i]
	}
	return string(ans)
}

func longestCommonSubstring(s string) string {
	N := len(s)
	if N <= 1 {
		return s
	}

	s2 := reverse(s)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}
	// dp
	var ans string
	for i := 0; i < N; i++ {
		if s[i] == s2[0] {
			dp[i][0] = 1
			ans = s[i : i+1]
		}
	}

	for j := 0; j < N; j++ {
		if s[0] == s2[j] {
			dp[0][j] = 1
			ans = s2[j : j+1]
		}
	}
	maxLen := 0
	for i := 1; i < N; i++ {
		for j := 1; j < N; j++ {
			if s[i] == s2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					ans = s[i+1-maxLen : i+1]
				}
			}
		}
	}
	return ans
}
