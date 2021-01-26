package main

/**
解题思路：问题转换为求相同数字组成窗口的左右下标；求解后，总数= right-left-1(因为左右下标会各自多移一位，所以是减一)
时间复杂度：logn
空间复杂度：1
*/
func search2(nums []int, target int) int {
	// 求窗口右下标
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] <= target { // target 在右区间
			i = m + 1
		} else { // target 在左区间
			j = m - 1
		}
	}
	right := i
	// 判断数据是否有 target
	if j >= 0 && nums[j] != target { // j 是窗口的右下标
		return 0
	}
	// 求窗口左下标
	i = 0
	for i <= j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else {
			j = m - 1
		}
	}
	left := j

	return right - left - 1
}
