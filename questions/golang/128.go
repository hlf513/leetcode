package main

import "fmt"

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

/**
 解题思路：哈希表
 	1. 构建哈希表
	2. 判断前后数字是否存在，存在+1
	3. 记录最大值
 时间复杂度：O(n)
 空间复杂度：O(n)
*/	
func longestConsecutive(nums []int) int {
	var hashmap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = i
	}

	var max int
	for v, _ := range hashmap {
		delete(hashmap, v)
		// pre
		pre := v - 1
		for {
			if _, ok := hashmap[pre]; ok {
				delete(hashmap, pre)
				pre--
			} else {
				break
			}
		}
		// next
		next := v + 1
		for {
			if _, ok := hashmap[next]; ok {
				delete(hashmap, next)
				next++
			} else {
				break
			}
		}
		if next-pre-1 > max {
			max = next - pre - 1
		}
	}

	return max
}
