package main

import "leetcode/common"

/**
解题思路：BFS，层序遍历，每层使用一维数组，组成二维数组
时间复杂度：N
空间复杂度：N
 */
func levelOrder(root *common.TreeNode) [][]int {
	var result [][]int
	
	if root == nil{
		return result
	}
	
	// 初始化队列
	var q []*common.TreeNode
	q = append(q,root)
	
	for len(q) != 0 {
		// 遍历当前层级
		var cur []int
		
		for _,item := range q {
			cur = append(cur,item.Val)	
			q = q[1:] // 保存后从队列中弹出
			
			if item.Left != nil{
				q = append(q,item.Left)	
			}
			if item.Right!= nil{
				q = append(q,item.Right)
			}
		}
		// 保存当前层级数据
		result = append(result,cur)
		
	}
	
	return result
}
