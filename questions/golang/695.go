package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		[]int{1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	// 递归
	//fmt.Println(maxAreaOfIsland(grid))
	// BFS
	fmt.Println(maxAreaOfIsland2(grid))
}

/**
 解题思路：递归
  	每个1要判断左、右、上、下
 时间复杂度：O(n^3)
 空间复杂度：O(1)
*/
func maxAreaOfIsland(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var res = 0.0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			var cnt = 0.0
			count(grid, i, j, &cnt, &res)
		}
	}

	return int(res)
}

func count(grid [][]int, i, j int, cnt, res *float64) {
	var m = len(grid)
	var n = len(grid[0])

	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] <= 0 {
		return
	}

	grid[i][j] = -1
	*cnt++
	*res = math.Max(*res, *cnt)

	dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	for _, dir := range dirs {
		count(grid, i+dir[0], j+dir[1], cnt, res)
	}
}

/**
 解题思路：BFS
 时间复杂度：O(n)
 空间复杂度：O(1)
*/
func maxAreaOfIsland2(grid [][]int) int {
	var m = len(grid)
	var n = len(grid[0])
	var res = 0.0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}

			var cnt = 0.0
			var q [][]int
			q = append(q, []int{i, j})
			grid[i][j] = -1

			for len(q) > 0 {
				item := q[0:1]
				q = q[1:]

				cnt++
				res = math.Max(res, cnt)

				dirs := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
				for _, dir := range dirs {
					x := item[0][0] + dir[0]
					y := item[0][1] + dir[1]

					if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] <= 0 {
						continue
					}

					grid[x][y] = -1
					q = append(q, []int{x, y})
				}
			}
		}
	}

	return int(res)
}
