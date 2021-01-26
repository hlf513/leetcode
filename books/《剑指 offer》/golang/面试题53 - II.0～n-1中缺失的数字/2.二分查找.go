package main

/**
解题思路：二分查找，因为索引和值是一样的，所以找 nums[i] != i，若相等，说明在右边
时间复杂度：logn
空间复杂度：1
*/
func missingNumber2(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		m := i + (j-i)/2
		if nums[m] == m { // 相等，说明在右边
			i = m + 1
		} else { // 不相等，说明在左边
			j = m - 1
		}
	}

	// 此时，i=右子数组的首位，j=左子数组的末位
	return i
}
