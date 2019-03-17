package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(nums, target))

}

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
