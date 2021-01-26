package main

/**
解题思路：遍历寻找最小数，优化：若前一个数大于当前数，说明是旋转后的最小数下标，可以直接返回
时间复杂度：N
空间复杂度：1
 */
func minArray(numbers []int) int {
	var min = numbers[0]
	for i:=1;i<len(numbers);i++{
		if numbers[i] < min{
			min = numbers[i]
		}
		// 找到旋转下标，直接返回
		if numbers[i] < numbers[i-1] {
			return min
		}
	}

	return min
}
