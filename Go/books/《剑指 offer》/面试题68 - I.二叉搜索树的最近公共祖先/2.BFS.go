package main

import "leetcode/common"

/**
解题思路：BFS
时间复杂度：N
空间复杂度：1
 */
func lowestCommonAncestor2(root, p, q *common.TreeNode) *common.TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val { // pq 在左子树
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val { // pq 在右子树
			root = root.Right
		} else { // 命中
			return root
		}
	}
	
	// 未命中
	return root
}
