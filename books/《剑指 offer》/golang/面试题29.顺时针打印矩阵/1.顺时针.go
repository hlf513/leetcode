package main


/**
解题思路： 模拟顺时针，找规律；定义四个坐标，每遍历一行/一列，减少相应的坐标，直到遍历结束
时间复杂度：N(行*列=元素个数)
空间复杂度：N(保存结果)
 */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	
	// 定义坐标
	var l,t = 0,0
	var r = len(matrix[0])-1
	var b = len(matrix)-1
	// 定义 result
	var result []int
	
	for l<=r && t <=b {
		// left to right 
		for i:=l;i<=r;i++ { // 注意 i
			result = append(result,matrix[t][i])		
		}
		t++ 
		if t > b { // 注意边界
			break	
		}
		// top to bottom
		for i:=t;i<=b;i++{
			result = append(result,matrix[i][r]) // 注意坐标哦
		}
		r--
		if r < l { // 注意越界
			break
		}
		// right to left
		for i:=r;i>=l;i--{ // 注意 i
			result = append(result,matrix[b][i])
		}
		b--
		if b < t { // 注意越界
			break
		}
		// bottom to top
		for i:=b;i>=t;i--{ // 注意 i
			result = append(result,matrix[i][l]) // 注意坐标哦
		}
		l++
		if l > r { // 注意越界
			break
		}
	}
	
	return result
}