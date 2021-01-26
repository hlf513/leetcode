package main

import "fmt"

func main() {
	fmt.Println(getLeastNumbers2([]int{3, 2, 1}, 2))
}

/**
解题思路： 排序+分治，只关注前 k 值
时间复杂度：N
空间复杂度：递归深度 logN~N
*/
func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	quickSelected(arr, 0, len(arr)-1, k)

	return arr[:k]
}

// partition 获取分区下标
func partition(nums []int, left, right int) int {
	i := left - 1
	pivot := nums[right]
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[right] = nums[right], nums[i+1]
	return i + 1
}

func quickSelected(nums []int, left, right, k int) {
	pos := partition(nums, left, right)
	num := pos - left + 1
	if k < num {
		quickSelected(nums, left, pos-1, k)
	} else if k > num {
		quickSelected(nums, pos+1, right, k-num)
	}
}
