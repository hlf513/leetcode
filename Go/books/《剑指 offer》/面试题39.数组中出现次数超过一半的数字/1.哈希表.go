package main

func main()  {
	
}

/**
解题思路：使用哈希表记录出现次数，一旦次数超过 n/2 则返回
时间复杂度：N
空间复杂度：N
*/
func majorityElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	var cache = make(map[int]int)
	for _,v := range nums{
		cache[v]++ // 先计算次数
		if num := cache[v];num > n/2{
			return v	
		}
	}

	return -1
}