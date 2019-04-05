package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}

/**
方法一：暴力循环 O(N^2)
方法二：Map O(N)
方法三：排序 O(NlogN)
方法四：分治 O(NlogN)
*/

// 方法二
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	var n = len(nums)
	for i := 0; i < n; i++ {
		m[nums[i]]++
		if m[nums[i]] > n/2 { // 超过一半则肯定出现最多
			return nums[i]
		}
	}
	var maxCount, res int
	for k, v := range m {
		if v > maxCount {
			res = k
		}
	}
	return res
}
