package main

import "fmt"

func main() {

	nums := []int{1, 3, 5, 4, 7}
	//nums := []int{2, 2, 2, 2, 2}
	//nums := []int{}

	fmt.Println(findLengthOfLCIS(nums))
}

/**
 解题思路：迭代+比较
 时间复杂度：O(n)
 空间复杂度：O(1)
*/
func findLengthOfLCIS(nums []int) int {
	max, c, n := 1, 0, len(nums)

	if n == 0 {
		return 0
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] > nums[i-1] {
			c++
			if max < c {
				max = c
			}
		} else {
			c = 1
		}
	}

	return max
}
