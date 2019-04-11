package main

import "fmt"

func main() {

	fmt.Println(hammingWeight(8))
}

/**
方法一：求模
方法二：位运算销最低位1
*/

// 方法一
func hammingWeight(num uint32) int {
	var ret int
	for num != 0 { // 不断/2,最后为0
		if num%2 != 0 { // num & 1
			ret++
		}
		num = num / 2 // num << 1
	}

	return ret
}

// 方法二
func hammingWeight2(num uint32) int {
	var ret int
	for num != 0 {
		ret++
		num = num & (num - 1)
	}
	return ret
}
