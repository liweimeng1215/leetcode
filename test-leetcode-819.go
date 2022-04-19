package main

import (
	"fmt"
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := map[string]int{}
	for _, str := range banned {
		bannedMap[str] = 1
	}

	ff := func(r rune) bool { return !unicode.IsLetter(r) }
	words := strings.FieldsFunc(paragraph, ff)
	wc := map[string]int{}
	max := 0
	ans := ""
	for _, w := range words {
		wl := strings.ToLower(w)
		if _, ok := bannedMap[wl]; ok {
			continue
		}
		wc[wl]++
		if wc[wl] > max {
			ans = wl
			max = wc[wl]
		}
	}
	return ans
}

func main() {
	ans := mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
	fmt.Printf("ans: %v\n", ans)
}
