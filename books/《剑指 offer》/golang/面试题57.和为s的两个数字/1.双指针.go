package main

/**
题目： 单调递增数组，找和为 s 的两个数字
解题思路：双指针，两数之和 < target 左指针+1，> target 右指针-1，相等则找到
时间复杂度：N
空间复杂度：1
*/
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
