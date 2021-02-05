package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(nums, target))

}

/**
 * 解题思路：二分法
 *		中间值 < 最右值：右边有序
 * 		中间值 > 最左值：左边有序
 * 时间复杂度：O(logn)
 * 空间复杂度：O(n)
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := int(left + (right-left)/2)
		//fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[right] { // 右边有序
			if nums[mid] < target && nums[right] >= target {
				// 右集合查找
				left = mid + 1
			} else {
				// 左集合查找
				right = mid - 1
			}
		} else {
			if nums[mid] > target && nums[left] <= target {
				// 左集合查找
				right = mid - 1
			} else {
				// 右集合查找
				left = mid + 1
			}
		}
	}

	return -1
}
