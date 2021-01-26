package main

import "fmt"

func main() {
	fmt.Println(myPow2(2.00, 4))
}

/**
方法一：库函数
方法二：暴力循环 O(N)
方法三：分治 O(logN)
*/

// 方法三
func myPow(x float64, n int) float64 {
	if x == 0 || n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}

	if n < 0 { // 负数
		return 1 / myPow(x, -n)
	}
	if n%2 > 0 { // 奇数
		return x * myPow(x, n-1)
	}

	return myPow(x*x, n/2)
}

// 方法三：非递归
func myPow2(x float64, n int) float64 {
	if n < 0 { // 负数
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n > 0 { // 正数
		if n&1 > 0 { // 奇偶性；==0 偶数
			pow *= x
		}
		x *= x
		n >>= 1 // n = n/2
	}

	return pow
}
