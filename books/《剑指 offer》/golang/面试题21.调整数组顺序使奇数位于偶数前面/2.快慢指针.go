package main

/**
解题思路： 快慢指针，从0开始，快指针先走，若是奇数和慢指针交换位置，同时慢指针移位，直到快指针走到数组尾部
时间复杂度：N
空间复杂度：1
*/
func exchange2(nums []int) []int {
	quick := 0
	slow := 0
	l := len(nums)
	
	for quick < l {
		if nums[quick] & 1 !=0 { // 快指针是奇数
			nums[quick],nums[slow] = nums[slow],nums[quick]
			slow++
		}
		quick++
	}
	
	return nums
}
