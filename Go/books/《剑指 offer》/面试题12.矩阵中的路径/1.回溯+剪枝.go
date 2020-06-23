package main

/**
题目：矩阵中是否存在字符串路径
解题思路：回溯+剪枝
时间复杂度：3^k (每部有三种选择，k 为 word 长度)
空间复杂度：k（递归深度，k 为 word 长度）
*/
// exist
func exist(board [][]byte, word string) bool {
	// m 行数，n 列数
	m, n := len(board), len(board[0])
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if search(row, col, 0, word, board) {
				return true
			}
		}
	}

	return false
}

// row 行坐标
// col 列坐标
// index word 的下标
// word 待找的字符串
// board 矩阵
func search(row, col, index int, word string, board [][]byte) bool {
	// 退出条件
	if index == len(word) { // 因为查下个字符时 index+1，所以会越界一位
		return true
	}
	if row < 0 || col < 0 || row == len(board) || col == len(board[0]) {
		return false
	}

	// 递归查找
	if board[row][col] == word[index] { // 当前字符命中
		// 假设命中
		tmp := board[row][col]
		board[row][col] = ' ' // 防止重复进入

		// 查下一个字符
		if search(row, col+1, index+1, word, board) ||
			search(row, col-1, index+1, word, board) ||
			search(row+1, col, index+1, word, board) ||
			search(row-1, col, index+1, word, board) {
			return true
		}

		// 未命中，恢复
		board[row][col] = tmp
	}

	return false
}
