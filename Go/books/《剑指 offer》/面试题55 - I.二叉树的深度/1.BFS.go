package main

import "leetcode/common"

/**
解题思路：BFS
时间复杂度：N
空间复杂度：N
 */
func maxDepth(root *common.TreeNode) int {
	// 边界条件
	if root == nil {
		return 0
	}

	// 队列
	var q []*common.TreeNode
	q = append(q, root)

	// 深度
	var level int

	// 遍历
	for len(q) > 0 {
		// 当前层级逻辑
		for _, node := range q {
			q = q[1:] // 出队
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		// 进入下个层级逻辑
		level++
	}

	return level
}