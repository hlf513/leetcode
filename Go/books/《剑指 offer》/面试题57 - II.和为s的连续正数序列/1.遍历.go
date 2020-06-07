package main

/**
题目：输出所有和为 target 的连续正整数序列（至少含有两个数）。
解题思路： 遍历，end = (target-1)/2
时间复杂度：target(根号target)
空间复杂度：1
*/
func findContinuousSequence(target int) [][]int {
	var result [][]int
	var sum int
	end := (target - 1) / 2
	for i := 1; i <= end; i++ { // 每个元素遍历一遍
		for j := i; j <= target; j++ { // 计算总和
			sum += j
			if sum > target {
				sum = 0
				break
			} else if sum == target {
				var item []int
				for k := i; k <= j; k++ {
					item = append(item, k)
				}
				result = append(result, item)
				sum = 0
				break
			}
		}
	}
	return result
}
