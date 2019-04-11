package main

import "fmt"

func main() {
	fmt.Println(mySqrt2(10))
}

/**
方法一：二分查找
方法二：牛顿迭代法
*/

// 方法一
func mySqrt(x int) int {
	left := 1
	right := x
	var mid int

	for left <= right {
		mid = left + (right-left)/2
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt > mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// 方法二
func mySqrt2(x int) int {
	var t int
	t = 0x5f3759df - (t >> 1)

	for !(t*t <= x && (t+1)*(t+1) > x) {
		t = (x/t + t) / 2
	}

	return t
}
