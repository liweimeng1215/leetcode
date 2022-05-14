// 7. 整数反转
// 反转int32类型整数，注意范围在 [-2^31, 2^31-1]
// 123 -> 321 -321 -> -123
// 将整数转为字符串，先反转字符串，然后进行处理
package main

import "fmt"

func reverseString(s string) string {
	N := len(s)
	ans := make([]byte, N)
	for i := 0; i < N; i++ {
		j := N - 1 - i
		ans[i] = s[j]
	}
	for s[0] == '0' {
		s = s[1:]
	}
	return string(ans)
}

func reverse(x int) int { // max int = 2^31-1 = 1<<31-1 ; min int = -2^31 = -1<<31
	if x == 0 {
		return 0
	}

	flag := false
	s := fmt.Sprintf("%v", x)
	if s[0] == '-' {
		flag = true
		s = s[1:]
	}

	rs := reverseString(s)
	n := len(rs)
	maxInt, minInt := 1<<31-1, -1<<31
	rx := 0
	if flag == false {
		for i := n - 1; i >= 0; i-- {
			curBit := int(s[i] - '0')
			if (maxInt-curBit)/10 < rx {
				return 0
			}
			rx = rx*10 + curBit
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			curBit := int(s[i] - '0')
			if (minInt+curBit)/10 > rx {
				return 0
			}
			rx = rx*10 - curBit
		}
	}
	return rx
}
