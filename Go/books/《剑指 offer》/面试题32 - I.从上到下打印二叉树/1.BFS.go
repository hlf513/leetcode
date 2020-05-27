package main

import (
	"leetcode/common"
)

/**
题目：从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
解题思路：BFS，层序遍历，
时间复杂度：N
空间复杂度：N
 */
func levelOrder(root *common.TreeNode) []int {
	var result []int
	
	if root == nil{
		return result
	}
	
	// 初始化队列
	var q []*common.TreeNode
	q = append(q,root) // 把根节点放入队列
	
	for len(q) != 0 {
		node := q[0] // 每次从队头取一个元素
		q = q[1:] // 更新队列

		result = append(result,node.Val) // 保存数值
		
		if node.Left != nil { // 若有左节点则把做节点放入队列
			q = append(q,node.Left)
		}
		if node.Right != nil{ // 若有右节点则把做节点放入队列
			q = append(q,node.Right)
		}
	}
	
	return result
}
