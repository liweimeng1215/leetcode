// LC5. 最长回文子串
// 字符串的最长回文串，中心扩展方法
package main

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	var ans = s[0:1]
	var maxLen int
	for i := 1; i <= (n-1)*2; i++ {
		var lp, rp int
		if i%2 == 0 {
			lp, rp = i/2-1, i/2+1
		} else {
			lp, rp = i/2, i/2+1
		}
		for 0 <= lp && rp <= n-1 && s[lp] == s[rp] {
			if rp+1-lp > maxLen {
				ans = s[lp : rp+1]
				maxLen = rp + 1 - lp
			}
			lp--
			rp++
		}
	}

	return ans
}
