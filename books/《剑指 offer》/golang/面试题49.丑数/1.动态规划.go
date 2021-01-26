package main

import "fmt"

/**
题目：求第 N 个丑数（能被2，3，5整除的数，1也是丑数）
解题思路：dp[i] = min(dp[a]*2,dp[b]*3,dp[c]*5) abc 对应最小值依次+1
时间复杂度：N
空间复杂度：N
*/
func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make(map[int]int)
	dp[0] = 1

	min := func(one, two int) int {
		if one < two {
			return one
		}

		return two
	}

	for i := 1; i < n; i++ {
		// 求丑数
		num := min(min(dp[a]*2, dp[b]*3), dp[c]*5)
		dp[i] = num
		// 下标自增
		if dp[a]*2 == num {
			a++
		}
		if dp[b]*3 == num {
			b++
		}
		if dp[c]*5 == num {
			c++
		}
	}

	return dp[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
