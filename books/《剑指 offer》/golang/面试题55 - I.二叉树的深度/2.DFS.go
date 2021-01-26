package main

import "leetcode/common"

/**
解题思路：DFS
时间复杂度：N
空间复杂度：N
 */
func maxDepth2(root *common.TreeNode) int {
	// 退出条件
	if root == nil {
		return 0
	}
	// 下沉
	left := maxDepth2(root.Left)
	right := maxDepth2(root.Right)

	// 判断左子树深度大，还是右子树深度大
	if left >= right {
		return left + 1
	}else {
		return right + 1
	}

}