package main

func main() {

}

/**
题目：找出所有逆序对的个数
解题思路： 枚举所有情况，判断是否是逆序对
时间复杂度：N^2
空间复杂度：1
*/
func reversePairs(nums []int) int {
	var result int
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] > nums[j] {
				result++
			}
		}
	}

	return result
}
