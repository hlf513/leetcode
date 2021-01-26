package main

import "fmt"

func main()  {
	
	s := ""
	fmt.Printf("%s",string(firstUniqChar(s)))
}

/**
解题思路：使用哈希表存储出现次数
时间复杂度：N
空间复杂度：N
 */
func firstUniqChar(s string) byte {
	var nums = make(map[byte]int)
	
	for _,v := range []byte(s) {
		nums[v]++
	}
	for _,v := range []byte(s) {
		if nums[v] == 1{
			return v
		}
	}
	
	return ' ' // 返回默认值
}

/**
解题思路：哈希表 + unicode 计数
 */
func firstUniqChar2(s string) byte {
	if s == ""{
		return ' '
	}
	var c [26]int
	for _,v := range s{
		c[v-'a']++ // 按照 unicode 计数
	}
	
	for _,v := range s{
		if c[v-'a'] == 1{
			return byte(v)
		}
	}
	
	return ' '
}
