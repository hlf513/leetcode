package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n1 = "3"
	var n2 = "56"

	fmt.Println(multiply(n1, n2))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 缓存中间值
	var tmp = make(map[int]int)
	c1 := len(num1)
	c2 := len(num2)
	for i := c2 - 1; i >= 0; i-- {
		for j := c1 - 1; j >= 0; j-- {
			n1 := string(num1[j])
			n2 := string(num2[i])
			cn1, _ := strconv.Atoi(n1)
			cn2, _ := strconv.Atoi(n2)

			// 计算下标(乘积最大长度为：c1+c2)
			tmp[c1+c2-2-i-j] += cn1 * cn2
		}
	}

	// 判断是否进位
	var res string
	var tc = len(tmp)
	//fmt.Println(tmp)
	for i := 0; i <= tc; i++ { // 注意这里的 <= tc，溢出一位是因为要判断进位
		tmp[i+1] += tmp[i] / 10
		res = strconv.Itoa(tmp[i]%10) + res
	}
	res = strings.TrimLeft(res, "0") // 去除开头0

	return res
}
