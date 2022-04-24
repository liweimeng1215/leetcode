/*
给定一个非空字符串S，其被N个‘-’分隔成N+1的子串，给定正整数K，要求除第一个子串外，其余的子串每K个字符组成新的子串，并用‘-’分隔。对于新组成的每一个子串，如果它含有的小写字母比大写字母多，则将这个子串的所有大写字母转换为小写字母；反之，如果它含有的大写字母比小写字母多，则将这个子串的所有小写字母转换为大写字母；大小写字母的数量相等时，不做转换。
输入描述:
输入为两行，第一行为参数K，第二行为字符串S。
输出描述:
输出转换后的字符串。
示例1
输入
3
12abc-abCABc-4aB@
输出
12abc-abc-ABC-4aB-@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每3个字符一组为abC、ABc、4aB、@，abC中小写字母较多，转换为abc，ABc中大写字母较多，转换为ABC，4aB中大小写字母都为1个，不做转换，@中没有字母，连起来即12abc-abc-ABC-4aB-@
示例2
输入
12
12abc-abCABc-4aB@
输出
12abc-abCABc4aB@
说明
子串为12abc、abCABc、4aB@，第一个子串保留，后面的子串每12个字符一组为abCABc4aB@，这个子串中大小写字母都为4个，不做转换，连起来即12abc-abCABc4aB@
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

const Debug = false

func toUpper(r []rune, p []int) string {
	for _, i := range p {
		r[i] -= 'a' - 'A'
	}
	return string(r)
}

func toLower(r []rune, p []int) string {
	for _, i := range p {
		r[i] += 'a' - 'A'
	}
	return string(r)
}

func min(x int, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func debug(format string, a ...any) {
	if Debug {
		fmt.Fprintf(os.Stdout, "[Debug] "+format, a...)
	}
}

func Translate(k int, s string) string {
	ff := func(r rune) bool { return r == '-' }
	strs := strings.FieldsFunc(s, ff)
	debug("split strs by '-': %v\n", strs)
	ans := strs[0]
	r := []rune(strings.Join(strs[1:], ""))
	debug("join strs without index 0: %v\n", string(r))

	for i, end := 0, k; i < end && i < len(r); i, end = end, end+k {
		tail := min(end, len(r))
		sub := r[i:tail]
		debug("split r by k: %v, start: %d, end: %d\n", string(sub), i, tail)
		var a1, a2 int
		for _, c := range sub {
			if 'A' <= c && c <= 'Z' {
				a1++
			} else if 'a' <= c && c <= 'z' {
				a2++
			}
		}

		if a1 > a2 {
			ans = ans + "-" + strings.ToUpper(string(sub))
		} else if a1 < a2 {
			ans = ans + "-" + strings.ToLower(string(sub))
		} else {
			ans = ans + "-" + string(sub)
		}
	}
	return ans
}

func main() {
	//var k int
	//var s string
	//fmt.Scanln(&k)
	//fmt.Scanln(&s)
	k := 12
	s := "12abc-abCABc-4aB@"
	fmt.Printf("k: %v, s: %v\n", k, s)
	ans := Translate(k, s)
	fmt.Printf("ans: %v\n", ans)
}
