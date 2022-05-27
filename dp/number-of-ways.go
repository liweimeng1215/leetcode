// LC1259. 不相交的握手
// dp解法
// 1. 选择 n 作为起点，dp[n]表示 n 个人的握手方案
// 2. n 点必须选择 1, 3, 5, ..., n-1 进行握手，否则会出现相交情况
// 3. dp[n] = sum {dp[0]*dp[n-2], dp[2]*dp[n-4], ..., dp[n-2]*dp[0]} <--> sum{ dp[2*j-2][n-2*j], j ∈ [1,2, ... , n/2] }
package main

import "fmt"

var d = false

func debug(format string, a ...interface{}) {
	if d {
		fmt.Printf(format, a...)
	}
}

func numberOfWays(numPeople int) int {
	var m int64 = 1000000007
	if numPeople%2 == 1 {
		return 0
	}
	dp := make([]int64, numPeople+1)
	dp[0] = 1

	for i := 2; i <= numPeople; i += 2 {
		var sum int64 = 0
		for j := 1; j <= i/2; j++ { // sum dp[2*j-2]*dp[i-2j]
			sum += (dp[2*j-2] % m) * (dp[i-2*j] % m) % m
		}
		dp[i] = sum
		debug("d[%v]: %v\n", i, dp[i])
	}
	return int(dp[numPeople] % m)
}
