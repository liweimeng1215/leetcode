// LC139. 单词拆分
// 遇到以下用例会超时，原因是出现了大量的重复计算，添加记忆map，避免重复计算
//
// "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
// ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"]
package main

import "fmt"

type Solution struct {
	WordDict []string
	memory   map[string]struct{}
}

var d = false

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

func (so *Solution) Search(s []byte) bool {
	if len(s) == 0 {
		return true
	}
	if _, ok := so.memory[string(s)]; ok {
		return false
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
	so.memory[string(s)] = struct{}{}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	So := Solution{WordDict: wordDict, memory: map[string]struct{}{}} // note: map need to initialize
	return So.Search([]byte(s))
}
