package main

/**
解题思路：二分查找思想，下标移位方式不同，每次都和「最右值」比较
时间复杂度：logn
空间复杂度：1
 */
func minArray2(nums []int) int {
	// 左右下标
	low := 0
	high := len(nums)-1
	
	for high > low{
		pivot := low + (high - low) / 2 // 取中位数	
		if nums[pivot] < nums[high]{ // 若中位数 < 最右数，则说明最小值在左侧
			high = pivot
		}else if nums[pivot] > nums[high]{ // 若中位数 > 最右数，则说明最小值在右侧
			low = pivot + 1
		}else{ // 若中位数 = 最右数，则最右数移位，重新开始比较
			high--
		}
	}
	return nums[low]
}
