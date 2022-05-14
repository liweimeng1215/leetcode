// LC14. 最长公共前缀
// 单词树解法，向单词树中插入字符串，增加数量与字符串长度的差即为公共前缀的长度
//
// 另一种更简单的解法：取出 strs[0] 与其它字符串比较，即可得出最长公共前缀。。
//
package main

import "fmt"

type Node struct {
	Nodes [26]*Node
}

type WordsTree struct {
	Root  *Node
	Count int
}

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func (wt *WordsTree) Insert(s string) int {
	p := wt.Root
	c1 := wt.Count
	for i := 0; i < len(s); i++ {
		id := s[i] - 'a'
		if p.Nodes[id] == nil {
			p.Nodes[id] = &Node{Nodes: [26]*Node{}}
			wt.Count++
		}
		p = p.Nodes[id]
	}
	return wt.Count - c1
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	lcpCount := len(strs[0])
	lcp := strs[0]

	tree := WordsTree{Root: &Node{Nodes: [26]*Node{}}}
	_ = tree.Insert(strs[0])

	for i := 1; i < n; i++ {
		count := tree.Insert(strs[i])
		curCpCount := len(strs[i]) - count
		debug("insert [%v] to words tree, insert c count: %v, lcpCount: %v\n", strs[i], curCpCount, lcpCount)
		if curCpCount < lcpCount {
			lcpCount = curCpCount
			lcp = strs[i][0:lcpCount]
			if lcpCount == 0 {
				return ""
			}
		}
	}
	return lcp
}
