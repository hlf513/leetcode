package main

import "fmt"

func main()  {
	fmt.Println(findRepeatNumber([]int{2,3,1,0,2,5,3}))	
}


/**
解题思路：哈希表，把 nums 依次放入哈希表，遍历到下一个是判断哈希表中是否存在
时间复杂度：N
空间复杂度：N
 */
func findRepeatNumber(nums []int) int {
	var cache = make(map[int]bool)
	l := len(nums)
	for i:=0;i<l;i++{
		if _,ok := cache[nums[i]]; ok {
			return nums[i]
		}
		cache[nums[i]] = true
	}
	
	return -1
}