//LC3. 无重复字符的最长子串
// 滑动窗口解法
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

func lengthOfLongestSubstring(s string) int {
	if strings.Compare(s, "") == 0 {
		return 0
	}
	window := map[byte]bool{}
	window[s[0]] = true
	ans, left := 1, 0
	for i := 1; i < len(s); i++ {
		if v, ok := window[s[i]]; !ok || v == false {
			window[s[i]] = true
			curLen := i - left + 1
			if curLen > ans {
				ans = curLen
			}
		} else {
			for window[s[i]] == true {
				window[s[left]] = false
				left++
			}
			window[s[i]] = true
		}
		debug("index: [%v], window: %v, windowDic: %v, ans: [%v]\n", i, s[left:i+1], window, ans)
	}
	return ans
}
