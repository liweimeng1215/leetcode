// 计算到字符串中各字符到某特定字符（存在多个）的最短距离

package main

import "fmt"

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func shortestToChar(s string, c byte) []int {
	places := make([]int, 0)
	for i, sc := range s {
		if byte(sc) == c {
			places = append(places, i)
		}
	}

	res := make([]int, len(s))
	if len(places) == 0 {
		return res
	} else if len(places) == 1 {
		for i := range s {
			res[i] = abs(places[0], i)
		}
		return res
	} else {
		start := 0
		for i := range places {
			if i == 0 {
				continue
			}
			end := (places[i] + places[i-1] + 1) / 2
			for j := start; j < end; j++ {
				res[j] = abs(j, places[i-1])
			}
			start = end
		}
		for i := start; i < len(s); i++ {
			res[i] = abs(i, places[len(places)-1])
		}
		return res
	}
}

func shortestToChar2(s string, c byte) []int {
	n := len(s)
	ans := make([]int, len(s))
	for i := range ans {
		ans[i] = n
	}
	j := -1
	for i := range s {
		if s[i] == c {
			j = i
		}
		if j != -1 {
			ans[i] = i - j
		}
	}

	j = -1
	for k := len(s) - 1; k >= 0; k-- {
		if s[k] == c {
			j = k
		}
		if j != -1 {
			ans[k] = min(ans[k], j-k)
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Printf("res: %v\n", shortestToChar("aaab", 'b'))
	fmt.Printf("res: %v\n", shortestToChar2("aaab", 'b'))
}
