package main

/**
题目：出现次数
解题思路：遍历，>目标值退出
时间复杂度：N
空间复杂度：1
*/
func search(nums []int, target int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
		if nums[i] > target {
			break
		}
	}
	return count
}
