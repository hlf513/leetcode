package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	fmt.Println(maxSlidingWindow(nums, k))
}

/**
方法一：大顶堆 (O(N logN))
方法二：双端队列 deque (O(N))
*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}
	// deque
	var window, res []int
	for i := 0; i < len(nums); i++ {
		// 划定范围
		if i >= k && len(window) > 0 && window[0] <= i-k { // i-k 表示需移除的最大下标
			window = window[1:]
		}
		// 保证window 第一个下标是最大值
		for len(window) > 0 && nums[window[len(window)-1]] <= nums[i] {
			window = window[:len(window)-1]
		}
		// 新增坐标（下一个迭代会处理溢出）
		window = append(window, i)
		// 保存最大值
		if i >= k-1 { // 过滤组建 window 过程时的最大值
			res = append(res, nums[window[0]])
		}
	}
	return res
}
