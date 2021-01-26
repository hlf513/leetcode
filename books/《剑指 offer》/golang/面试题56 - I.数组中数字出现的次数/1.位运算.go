package main

/**
题目：找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
解题思路：位运算；利用「任何数和自己异或都为0」& 分组异或
时间复杂度：N
空间复杂度：1
*/
func singleNumbers(nums []int) []int {
	// 获取所有数字的异或结果（得到两个出现一次数字的异或值）
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	// 在这个异或结果中找到任意为1的位
	mask := result & -result // 获取最低位的1
	// 分组异或
	var one, two int
	for i := 0; i < len(nums); i++ {
		// 分组
		if mask&nums[i] == 0 {
			// 异或后获取出现一次的数字
			one ^= nums[i]
		} else {
			two ^= nums[i]
		}
	}

	return []int{one, two}
}
