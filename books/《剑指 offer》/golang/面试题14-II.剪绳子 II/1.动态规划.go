package main

/**
题目：给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。 答案取模1000000007
解题思路：动态规划，问题抽象为求整数n分解后乘积最大,写例子找规律，只有2、3不会再次切分，所以 dp[i]=dp[i-3]*3
时间复杂度：N
空间复杂度：N
*/
func cuttingRope(n int) int {
	// 写例子得出的规律
	if n <= 3 {
		return n - 1
	}
	var dp = make(map[int]int) // dp[n] 表示 n 的最大乘积
	dp[2] = 2
	dp[3] = 3
	dp[4] = 4
	for i := 5; i <= n; i++ {
		dp[i] = 3 * dp[i-3] % 1000000007
	}

	return dp[n]
}
