package main

import (
	"fmt"
)

func main() {

	var nums = []int{2, 7, 11, 15}
	var target = 9

	fmt.Println(twoSum(nums, target))
}

/**
 * 解题思路：Hash，把遍历过的值放入 hash 中，若 target 存在，则直接返回
 * 时间复杂度：O(n)
 * 空间复杂度：O(2n)
*/
func twoSum(nums []int, target int) []int {
	var res []int
	var hash = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		search := target - nums[i]
		if _, ok := hash[search]; ok {
			return []int{search, nums[i]}
		}

		hash[nums[i]] = 1
	}

	return res
}
