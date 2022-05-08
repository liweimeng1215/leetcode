// LC415. 字符串相加
package main

import (
	"fmt"
	"strings"
)

var d = true

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func reverse(ans []byte) {
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
}

func addStrings(num1 string, num2 string) string {
	if strings.Compare(num1, "0") == 0 {
		return num2
	}
	if strings.Compare(num2, "0") == 0 {
		return num1
	}

	var i, j = len(num1) - 1, len(num2) - 1
	var ans = make([]byte, 0)
	addOne := 0
	for i >= 0 && j >= 0 {
		cur := (int(num1[i]-'0') + int(num2[j]-'0') + addOne) % 10
		ans = append(ans, byte('0'+cur))
		addOne = (int(num1[i]-'0') + int(num2[j]-'0') + addOne) / 10
		i--
		j--
	}

	func1 := func(a string, k int, ans []byte) []byte {
		for k != -1 {
			cur := (int(a[k]-'0') + addOne) % 10
			ans = append(ans, byte('0'+cur))
			addOne = (int(a[k]-'0') + addOne) / 10
			k--
		}
		return ans
	}

	if i >= 0 {
		ans = func1(num1, i, ans)
	}
	if j >= 0 {
		ans = func1(num2, j, ans)
	}
	if addOne > 0 {
		ans = append(ans, '1')
	}
	reverse(ans)

	return string(ans)

}
