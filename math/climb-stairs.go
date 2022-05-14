// LC70. 爬楼梯
// 经典斐波那契数列问题
// f(n) = f(n-1) + f(n-2)
package main

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	pre, cur := 1, 1
	for i := 2; i <= n; i++ {
		temp := cur
		cur = pre + cur
		pre = temp
	}
	return cur
}
