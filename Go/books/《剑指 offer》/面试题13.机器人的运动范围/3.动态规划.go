
type point3 struct {
	row int
	col int
}

var result map[point3]bool

/**
解题思路：因为只能向右、向下移动， 从下往上思考，一个符合条件的格子，只能从它的上、左走过来，那只要判断所有格子的左、上是否符合条件即可，注意判重
时间复杂度：M*N
空间复杂度：M*N
*/
func movingCount3(m int, n int, k int) int {
	result = make(map[point3]bool)
	result[point3{0, 0}] = true
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			// 判断这个格子的左、上是否在符合条件的集合中，存在则判断这个格子是否符合条件
			if (result[point3{row,col-1}] == true || result[point3{row-1,col}] == true ) && add3(row,col) <=k{
					result[point3{row,col}] = true // 符合条件放入集合中
			}
		}
	}

	return len(result)
}

// add3 按位相加
func add3(one, two int) int {
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
