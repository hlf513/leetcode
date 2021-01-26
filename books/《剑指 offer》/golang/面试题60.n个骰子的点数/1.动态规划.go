package main

import "math"

/**
题目：把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
解题思路：动态规划
	- base case：当投掷完 i 枚骰子后，各个点数出现的次数 j
	- 状态：dp[i][j]，表示投掷完 i 枚骰子后，点数 j 的出现次数
	- 状态转移方程：dp[n][j] = {dp[n-1][j-1]+dp[n-2][j-2]+...+dp[n-6][j-6]}
时间复杂度：N
空间复杂度：N
*/
func twoSum(n int) []float64 {
	var dp = make(map[int]int)
	// 初始化第一个骰子的点数出现次数
	for i := 1; i <= 6; i++ {
		dp[i] = 1
	}
	// 计算第 n 个骰子的点数出现次数
	for i := 2; i <= n; i++ {
		// 计算每个点数出现的次数
		for j := 6 * i; j >= i; j-- {
			dp[j] = 0
			// 每个点数 = 之前6个点数之和
			for cur := 1; cur <= 6; cur++ {
				if j-cur < i-1 {
					break
				}
				dp[j] += dp[j-cur]
			}
		}
	}
	// 计算所有点数的总次数
	all := math.Pow(6.0, float64(n))
	// 计算概率
	var result []float64
	for i := n; i <= 6*n; i++ {
		result = append(result, float64(dp[i])/all)
	}

	return result
}
