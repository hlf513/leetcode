package main

import (
	"fmt"
)

func main() {

	height := []int{0, 1, 0, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//height := []int{2, 1, 0, 1, 3}
	//height := []int{3, 2, 1, 2, 1}

	fmt.Println(trap(height))
}

// 数组法
/**
从两边向中间扫描，
找出较小值，
如果较小值是left指向的值，则从左向右扫描，
如果较小值是right指向的值，则从右向左扫描，
若遇到的值比当较小值小，则将差值存入结果，
如遇到的值大，则重新确定新的窗口范围，
以此类推直至left和right指针重合
*/
func trap(height []int) int {
	var left = 0
	var right = len(height) - 1
	var res = 0

	for left < right {
		// 找最小值
		var min int
		if height[left] >= height[right] {
			// right 是最小值，则当前值从右往左找
			min = height[right]
			// 当前值小于最小值，累积差值
			right = right - 1
			for left < right && height[right] < min {
				res = res + min - height[right]
				right = right - 1
			}
		} else {
			// left 是最小值，则当前值从左往右找
			min = height[left]
			// 当前值小于最小值，累积差值
			left = left + 1
			for left < right && height[left] < min {
				res = res + min - height[left]
				left = left + 1
			}
		}
	}

	return res
}
