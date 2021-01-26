package main

import "leetcode/common"

/**
解题思路：DFS
时间复杂度：N
空间复杂度：N
 */
func lowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	// 退出条件
	// 找到p,q，或者root 为空
	if root == nil || p == root || q == root {
		return root
	}

	// 在左子树查找 p,q
	left := lowestCommonAncestor(root.Left,p,q)
	// 在右子树查找 p,q
	right := lowestCommonAncestor(root.Right,p,q)

	if left == nil {
		// 左子树没有找到,说明在右子树中
		return right
	}
	if right == nil {
		// 右子树没有找到,说明在左子树中
		return left
	}
	// 左右子树都没有找到，返回当前的根结点
	return root
}