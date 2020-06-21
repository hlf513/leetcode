package main

import (
	"math"
	"sort"
)

/**
解题思路：把 nums 排序，然后取 n/2
时间复杂度：取决于排序算法(nlogn/n)
空间复杂度：取决于排序算法(logn/1)
 */
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	index := math.Ceil(float64(len(nums)/2))
	return nums[int(index)]
}
