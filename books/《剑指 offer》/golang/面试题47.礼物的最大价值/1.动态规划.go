package main

import "fmt"

/**
题目：在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
解题思路：动态规划，画出状态转移表（二维数组）
时间复杂度：MN
空间复杂度：1
*/
func maxValue(grid [][]int) int {
	row := len(grid)    // 行数
	col := len(grid[0]) // 列数
	// 计算第一行数值
	for i := 1; i < col; i++ {
		grid[0][i] += grid[0][i-1]
	}
	// 计算第一列数值
	for i := 1; i < row; i++ {
		grid[i][0] += grid[i-1][0]
	}

	// 计算其他数值
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			// 取上左最大值与自己相加
			grid[i][j] += max(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[row-1][col-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var g = [][]int{
		{1, 2, 5},
		{3, 2, 1},
	}
	fmt.Println(maxValue(g))
}
