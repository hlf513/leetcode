package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums = []int{-1, 0, 1, 2, -1, 1, -2}

	fmt.Println(threeSum(nums))
}

/**
 * 解题法：排序+三指针+剪枝
 * 解题思路：排序 + 使用三个指针 + 剪枝
 * 		1. 固定一个a，用左右指针找和为-a的两个数
 *		2. 注意三个指针都要去重
 * 		3. 注意边界
 * 时间复杂度：O(nlogn)
 * 空间复杂度：O(n)
*/
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	// 边界判断
	if n == 0 || nums[0] > 0 || nums[n-1] < 0 {
		return res
	}

	for i := 0; i < n; i++ {
		// 剪枝；正数永不满足条件
		if nums[i] > 0 {
			return res
		}
		// 重复跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0 - nums[i]
		l := i + 1
		r := n - 1

		for l < r {
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				// 左右指针移动（去重）
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] > target {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
	}

	return res
}
