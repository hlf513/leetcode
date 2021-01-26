package main

import "leetcode/common"

func main() {
	
}

/**
解题思路：DFS + 保存路径
时间复杂度：N
空间复杂度：N
 */
func pathSum(root *common.TreeNode, sum int) [][]int {
	result = [][]int{} // 初始化 result（全局变量，运行每个 case 前需要初始化）
	recursion(root,sum,[]int{})
	return result
}

var result [][]int

func recursion(root *common.TreeNode, sum int, cur []int)  {
	// 退出条件
	if root == nil{
		return
	}
	
	// 记录遍历路径
	cur = append(cur,root.Val)
	// 叶子节点，判断是否满足条件
	if root.Left == nil && root.Right == nil && sum == root.Val { // 满足条件
		// 重新 copy 一份路径，避免后续被污染
		path := make([]int,len(cur))
		copy(path,cur)
		result = append(result,path)	
	}
	
	// 下沉
	recursion(root.Left,sum-root.Val,cur)
	recursion(root.Right,sum-root.Val,cur)
}