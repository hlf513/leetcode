package main

type point2 struct {
	row int
	col int
}

var visited map[point2]bool

/**
解题思路：因为只能向右、向下移动，问题抽象为 DFS，子节点为向下、向右，注意防止重复计数，递归返回值为：右边数量+下面数量+1
时间复杂度：M*N
空间复杂度：M*N
*/
func movingCount2(m int, n int, k int) int {
	visited = make(map[point2]bool) // 在这里初始化，因为多 case 场景下，每个 case 需要重新初始化 visited
	return dfs(m, n, 0, 0, k)
}

// dfs 返回符合条件的数量
// m 行数
// n 列数
// row 当前第几行
// col 当前第几列
// k 数位和不能超过的数
func dfs(m, n, row, col, k int) int {
	if visited[point2{row, col}] == false && row < m && col < n && add2(row, col) <= k {
		visited[point2{row, col}] = true // 注意增加判重
		// 右+下+1（本身）
		return dfs(m, n, row, col+1, k) + dfs(m, n, row+1, col, k) + 1
	}
	// 退出条件
	return 0
}

// add2 按位相加
func add2(one, two int) int {
	var result int
	// 135
	for one != 0 {
		result += one % 10 // 5,3,1
		one = one / 10     // 13,1,0
	}
	// 10
	for two != 0 {
		result += two % 10
		two = two / 10
	}

	return result
}
