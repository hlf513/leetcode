package main

import (
	"sort"
	"strconv"
)

func main()  {
	
}

// 声明新类型，并定义排序规则
type Slice []int
func (p Slice) Len() int { return len(p) }
func (p Slice) Less(i, j int) bool {
	is := strconv.Itoa(p[i])
	js := strconv.Itoa(p[j])

	s1 := is + js
	s2 := js + is

	if s1 < s2 {
		return true
	}

	return false
}
func (p Slice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

/**
解题思路：按照字符串排序，然后拼接字符串；使用内置排序接口
时间复杂度：logn
空间复杂度：N
 */
func minNumber(nums []int) string {
	var nums2 = Slice(nums)
	sort.Sort(nums2)

	var ret string
	for _, num := range nums2 {
		ret += strconv.Itoa(num)
	}
	return ret
}
