// 7. 整数反转
// 反转int32类型整数，注意范围在 [-2^31, 2^31-1]
// 123 -> 321 -321 -> -123
// 实际上并不需要将整数转为字符串
// 直接将最后一位弹出
//	cur := x % 10
// 再将cur添加到结果 rx 的最后一位中
// rx = rx*10 + cur
package main

func reverse(x int) int { // max int = 2^31-1 = 1<<31-1 ; min int = -2^31 = -1<<31
	rx := 0
	maxInt := 1<<31 - 1
	minInt := -1 << 31
	for x != 0 {
		cur := x % 10
		if rx > 0 && (maxInt-cur)/10 < rx { // 注意范围，避免溢出
			return 0
		}
		if rx < 0 && (minInt-cur)/10 > rx {
			return 0
		}
		rx = rx*10 + cur
		x /= 10
	}
	return rx
}
