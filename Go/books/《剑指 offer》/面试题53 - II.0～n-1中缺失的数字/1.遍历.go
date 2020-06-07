package main

func main() {

}

/**
解题思路： 一次遍历，找到不存在的返回
时间复杂度：N
空间复杂度：1
*/
func missingNumber(nums []int) int {
	var n = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != n {
			return n
		}
		n++
	}
	return n
}
