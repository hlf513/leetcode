package main

import "fmt"

/**
题目：从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字
解题思路：判重 + max-min < 5；0跳过
时间复杂度：N
空间复杂度：1
*/
func isStraight(nums []int) bool {
	var max, min = 0, 14 // 最值初始化小技巧
	var repeat = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 大小王跳过
		if nums[i] == 0 {
			continue
		}
		// 判重
		if _, ok := repeat[nums[i]]; ok {
			return false
		}
		repeat[nums[i]] = 1
		// 选择最值
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max-min < 5
}

func main() {
	fmt.Println(isStraight([]int{0, 0, 1, 2, 5}))
}
