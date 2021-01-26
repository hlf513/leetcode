package main

import "fmt"

func main()  {
	
	s := "abcabcbb"
	s = "bbbb"
	// s = ""
	s = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
	
}

/**
解题思路： 滑动窗口（哈希表+双指针），哈希表存储字符出现的次数（用于去重），左指针指向字符的开始下标，右指针不断后移直到遇到重复字符；一旦右指针遇到重复字符，则计算当前子串最大长度；然后左指针右移，直到左指针到最后一位。
时间复杂度：N
空间复杂度：字符集
*/
func lengthOfLongestSubstring(s string) int {
	var result int // 最大长度
	var nums = make(map[byte]int) // 哈希表
	var left,right = 0,-1 // 左右指针，注意右指针从-1开始（因为还未开始移动）

	n := len(s)
	for left = 0; left < n; left++ {
		// 右指针开始右移，直到遇到重复字符
		for right+1<n && nums[s[right+1]] == 0 {
			nums[s[right+1]]++
			right++	
		}
		// 计算最大长度
		l := right-left+1
		if result < l {
			result = l
		}
		// 左指针右移,从哈希表中删除对应的计数
		delete(nums,s[left])
	}
	
	return result
}