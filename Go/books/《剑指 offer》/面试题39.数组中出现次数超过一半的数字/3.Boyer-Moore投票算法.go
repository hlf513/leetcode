package main

/**
解题思路：投票，第一个元素为第一个候选者，遍历每个元素，若票数=0则替换候选者，然后若元素=候选者则票数+1，反之减一；遍历结束时的候选者就是众数
时间复杂度：N
空间复杂度：1
 */
func majorityElement3(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	var candidate = nums[0]
	var count = 1
	for i:=1;i<l;i++{
		if count == 0 {
			candidate = nums[i]
			// 注意以下两行，优化执行效率（跳过没有必要比较）
			count = 1
			continue
		}
		if nums[i] == candidate {
			count++
		}else{
			count--
		}
	}

	return candidate
}
