package main

import "leetcode/common"

/**
解题思路：递归，直接判断两颗树是都对称，current.level 判断是否对称
时间复杂度：N
空间复杂度：N
 */
func isSymmetric(root *common.TreeNode) bool{
	return isMirror(root,root)
}

// isMirror 判断是否对称
func isMirror(root1,root2 *common.TreeNode) bool {
	// 退出条件
	if root1 == nil && root2 == nil { 
		return true
	}
	if root1 == nil || root2 == nil { 
		return false
	}
	
	// 判断是否对称
	return root1.Val == root2.Val &&
		isMirror(root1.Left,root2.Right) &&
		isMirror(root1.Right,root2.Left)
}