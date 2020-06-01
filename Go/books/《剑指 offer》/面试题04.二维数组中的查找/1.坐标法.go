package main

/**
题目：有序的矩阵中查找
解题思路： 从第一行最后一列开始查找，若大于则移动列，若小于则移动行
时间复杂度：m+n (行数+列数)
空间复杂度：1
 */
func findNumberIn2DArray(matrix [][]int, target int) bool {
	// 边界条件
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	
	// 计算列数
	columns := len(matrix[0])
	// 初始化第一个坐标(第一行row最后一列column)
	var row = 0
	var column = columns -1

	// 开始遍历
	for row < rows && column >= 0 {
		num :=  matrix[row][column]
		// 找到
		if num == target {
			return true
		}
		// 未找到
		if num > target{ // 大于则移动列
			column--
		}else{ // 小于移动行
			row++
		}
	}

	return false
}