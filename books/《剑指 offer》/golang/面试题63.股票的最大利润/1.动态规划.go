package main

/**
题目：假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
解题思路：动态规划，profit = max(历史最高利润, 当天价格 - 之前历史最低价格)
时间复杂度：N
空间复杂度：1
*/
func maxProfit(prices []int) int {
	// 边界条件
	if len(prices) == 0 {
		return 0
	}
	// 计算利润
	var cost = prices[0]
	var profit = 0
	for i := 0; i < len(prices); i++ {
		cost = min(cost, prices[i])          // 历史最小花销
		profit = max(profit, prices[i]-cost) // 历史最大利润
	}

	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
