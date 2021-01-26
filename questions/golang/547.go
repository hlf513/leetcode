package main

import "fmt"

func main() {

	//M := [][]int{
	//	{1, 1, 0},
	//	{1, 1, 0},
	//	{0, 0, 1},
	//}

	M := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 1, 1},
	}

	fmt.Println(findCircleNum(M))
}

func findCircleNum(M [][]int) int {
	var v = make(map[int]int)
	var res int
	for i := 0; i < len(M); i++ {
		// 判断是否已在朋友圈
		if _, ok := v[i]; ok {
			continue
		}
		// 遍历朋友圈
		helper(M, i, v)
		res++
	}

	return res
}

func helper(M [][]int, i int, v map[int]int) {
	// 退出条件
	if _, ok := v[i]; ok {
		return
	}
	// 处理逻辑
	v[i] = 1

	// 下沉
	for j := 0; j < len(M); j++ {
		_, ok := v[j]
		if M[i][j] == 0 || ok {
			continue
		}
		helper(M, j, v)
	}
}
