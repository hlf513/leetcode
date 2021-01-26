package main

/**
解题思路：首尾指针，若首指针为偶数，尾指针为奇数时，交换值；其他情况指针移位；直到双指针相遇
时间复杂度：N
空间复杂度：1
 */
func exchange(nums []int) []int {
	head := 0
	tail := len(nums)-1
	
	for head < tail {
		if nums[head] & 1 == 0  && nums[tail] & 1 != 0 { // 首指针为偶数，尾指针为奇数
			nums[head],nums[tail] = nums[tail],nums[head]
		}else {
			if nums[head] & 1 != 0 { // 首指针为奇数
				head++
			}
			if nums[tail] & 1 == 0{ // 尾指针为偶数
				tail--
			}
		}
	}
	
	return nums
}
