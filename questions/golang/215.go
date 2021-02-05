package main

import "fmt"

func main() {
	//nums := []int{3, 2, 7, 3, 1, 2, 4, 5, 5, 6}
	//nums := []int{3, 1, 6, 2, 4, 5}
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2

	fmt.Println(findKthLargest(nums, k))
}

/**
 解题思路：排序
	1. 快排（从大到小）
	2. 取下标 k-1 的值
 时间复杂度：O(nlogn~n^2)
 空间复杂度：O(n)
*/
func findKthLargest(nums []int, k int) int {
	qSort(nums)
	return nums[k-1]
}

func qSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mid := nums[0]
	head, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] < mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}

	qSort(nums[:head])
	qSort(nums[head+1:])
}
