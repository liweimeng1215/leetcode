// LC139. 单词拆分
package main

import "fmt"

type Solution struct {
	WordDict []string
}

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func compare(s1, s2 []byte) bool {
	debug("s1: %s, s2: %s\n", s1, s2)
	for i := range s2 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func (so Solution) Search(s []byte) bool {
	if len(s) == 0 {
		return true
	}

	ans := false
	for _, w := range so.WordDict {
		n := len(w)
		if len(s) < n {
			continue
		}
		comAns := compare(s, []byte(w))
		if comAns == false {
			continue
		}
		ans = so.Search(s[n:])
		if ans == true {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	So := Solution{WordDict: wordDict}
	return So.Search([]byte(s))
}
