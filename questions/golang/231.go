package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo2(16))
}

/**
方法一：mod/2
方法二：位运算（2的幂的二进制只有一个1）
*/

// 方法一
func isPowerOfTwo(n int) bool {
	// 注意这里的边界情况
	if n <= 0 {
		return false
	}

	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

// 方法二
func isPowerOfTwo2(n int) bool {
	// 注意这里的边界情况
	if n <= 0 {
		return false
	}

	if n&(n-1) == 0 {
		return true
	}
	return false
}
