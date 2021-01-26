package main

func main() {
	constructArr([]int{1, 2, 3, 4, 5})

}

/**
题目：给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。
解题思路：画出矩阵，分左三角( 先赋值为 b[i])、右三角( 计算结果后，再和 b[i]相乘)
时间复杂度：N
空间复杂度：1
*/
func constructArr(a []int) []int {
	// 边界条件
	if len(a) == 0 {
		return []int{}
	}
	// 初始化 b0 = 1
	var b = []int{1}
	// 计算左三角
	for i := 1; i < len(a); i++ {
		b = append(b, b[i-1]*a[i-1])
	}
	// 计算右三角
	var tmp = 1
	for i := len(a) - 2; i >= 0; i-- { // 从倒数第二个开始
		tmp *= a[i+1]
		b[i] *= tmp
	}

	return b
}
