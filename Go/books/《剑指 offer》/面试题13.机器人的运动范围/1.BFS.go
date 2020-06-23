package main

type point struct {
	row int
	col int
}

/**
题目：地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
解题思路：因为只能向右、向下移动，问题抽象为 BFS，子节点为向下、向右，注意防止重复计数
时间复杂度：M*N
空间复杂度：M*N
*/
func movingCount(m int, n int, k int) int {
	var visited = make(map[point]bool)
	var queue []point
	var result int
	// 初始化队列
	queue = append(queue, point{0, 0})
	// BFS
	for len(queue) > 0 {
		// 出队
		tmp := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// 判重+符合条件后计数
		if visited[tmp] == false && tmp.row < m && tmp.col < n && add(tmp.row, tmp.col) <= k {
			visited[tmp] = true
			result++
			// 把右、下放入队列
			queue = append(queue, point{tmp.row, tmp.col + 1}) // 右
			queue = append(queue, point{tmp.row + 1, tmp.col}) // 下
		}
	}

	return result
}

// add 按位相加
func add(one, two int) int {
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
