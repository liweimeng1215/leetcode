//LC3. 无重复字符的最长子串
//动态规划的思想
// 记录前一字符为结尾的无重复字符最长子串 f(n-1)
// f(n-1) = f(n)+s[i] 					if s[i] not in f(n)
//			f(n)[j+1:]+s[i] 			if s[i] in f(n), j is index of s[i] in f(n)
//
package main

import (
	"fmt"
	"strings"
)

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func indexOfStr(s []byte, t byte) int {
	for i, c := range s {
		if c == t {
			return i
		}
	}
	return -1
}

func lengthOfLongestSubstring(s string) int {
	if strings.Compare(s, "") == 0 {
		return 0
	}

	var old = make([]byte, 1)
	old[0] = s[0]
	var ans = 1
	for i := 1; i < len(s); i++ {
		j := indexOfStr(old, s[i])
		if j == -1 {
			old = append(old, s[i])
			if len(old) > ans {
				ans = len(old)
			}
		} else {
			if j == len(old)-1 {
				old = []byte{s[i]}
			} else {
				old = append(old[j+1:], s[i])
			}
		}
		debug("old: %s, ans: %v\n", old, ans)
	}
	return ans
}
