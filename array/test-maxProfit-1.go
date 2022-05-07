// LC121. 买卖股票的最佳时机
package main

func maxProfit(prices []int) int {
	var lowPrice = prices[0]
	var ans = 0
	for i := 1; i < len(prices); i++ {
		var profit = prices[i] - lowPrice
		if profit > 0 && profit > ans {
			ans = profit
		}
		if prices[i] < lowPrice {
			lowPrice = prices[i]
		}
	}
	return ans
}
