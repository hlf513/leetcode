package main

import "fmt"

func main() {
	p := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(p))
}

/**
方法一：DFS 找到所有情况 O(2^n)
方法二：贪心算法 O(N)
方法三：动态规划 O(N)
*/

// 方法二
func maxProfit(prices []int) int {
	n := len(prices)
	res := 0
	for i := 0; i < n; i++ {
		if i+1 < n && prices[i+1] > prices[i] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
