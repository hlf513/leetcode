package main

import "leetcode/common"

/**
解题思路：DFS，通过深度遍历
时间复杂度：N
空间复杂度：N
 */
func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	if p.Val > root.Val && q.Val > root.Val { // pq 在右子树
		return lowestCommonAncestor(root.Right,p,q)
	}else if p.Val < root.Val && q.Val < root.Val { // pq 在左子树
		return lowestCommonAncestor(root.Left,p,q)
	}else{ // 命中
		return root
	}
}
