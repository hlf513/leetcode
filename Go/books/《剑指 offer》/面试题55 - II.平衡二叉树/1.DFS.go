package main

import "leetcode/common"

/**
解题思路：通过 DFS 获取子树的高度后，判断是否是平衡
时间复杂度：N
空间复杂度：N
 */
func isBalanced(root *common.TreeNode) bool {
	// 边界条件
	if root == nil {
		return true
	}
	
	// -1 表示不是平衡树
	return dfs(root) != -1
}

// dfs 获取节点的子树的深度后，判断是否平衡
func dfs(root *common.TreeNode) int {
	// 退出条件（这里的0表示当前节点的深度）
	if root == nil {
		return 0
	}

	// 下沉
	left := dfs(root.Left)
	right := dfs(root.Right)

	// 判断是否平衡，-1 表示不是平衡树
	if abs(left-right) > 1 {
		return -1
	}
	if left == - 1 || right == -1 {
		return -1
	}

	// 返回子树的深度
	return max(left,right) + 1
}

func abs(i int) int {
	if i > 0 {
		return i
	} else {
		return -i
	}
}

func max(a,b int) int{
	if a >= b {
		return a
	}
	return b
}
