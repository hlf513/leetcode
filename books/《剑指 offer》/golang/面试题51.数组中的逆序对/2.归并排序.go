package main

/**
题目：获取逆序对总数
解题思路： 利用归并排序的合并、稳定排序特性来计算逆序对；合并时，若左<右，则右移位的数量就是左边这个数字的逆序对总数。
时间复杂度：nlogn
空间复杂度：1
*/
func reversePairs2(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, start, end int) int {
	// 退出条件
	if start >= end {
		return 0
	}
	// 找中位
	mid := start + (end-start)/2
	// 逆序对计数
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	// 排序 && 计数
	i, j := start, mid+1
	var sorted = []int{}
	for i <= mid && j <= end {
		if nums[i] <= nums[j] { // 从小到大；注意是<=
			sorted = append(sorted, nums[i])
			count += j - (mid + 1) // 不包含当前指针
			i++                    // 左指针移位
		} else {
			sorted = append(sorted, nums[j])
			j++ // 右指针移位
		}
	}
	// 处理左边剩余数据
	for ; i <= mid; i++ {
		sorted = append(sorted, nums[i])
		count += end - (mid + 1) + 1
	}
	// 处理右边剩余数据
	for ; j <= end; j++ {
		sorted = append(sorted, nums[j])
	}
	// 写入原数组
	for i := start; i <= end; i++ {
		nums[i] = sorted[i-start]
	}

	// 返回逆序对数
	return count
}
