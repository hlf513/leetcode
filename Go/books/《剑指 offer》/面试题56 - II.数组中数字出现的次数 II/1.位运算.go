package main

/**
题目：在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。
解题思路： 位运算，每个数字按位统计，统计每位1的个数，把最终得到的数 % 3 = 出现一次的数字
时间复杂度：N
空间复杂度：1
*/
func singleNumber(nums []int) int {
	// 边界条件
	if len(nums) == 1 {
		return nums[0]
	}
	// 统计位1的个数
	var counts = make([]int, 32) // 假设是 32位 int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < 32; j++ {
			counts[j] += nums[i] & 1 // 同样是1的话，计数+1
			nums[i] >>= 1            // nums[i]右移一位
		}
	}
	// 按位对 3 取余后，按位或
	res, m := 0, 3
	for i := 0; i < 32; i++ {
		res = res << 1          // 左移一位
		res |= counts[31-i] % m // 恢复第 i 位的值到 res
	}
	return res
}

func main() {
	// singleNumber([]int{3, 5, 3, 3})

}
