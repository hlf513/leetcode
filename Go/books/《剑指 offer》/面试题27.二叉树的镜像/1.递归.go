package main

import "leetcode/common"

func main()  {
	
}

/**
解题思路：递归，current level：直接交换左右节点
时间复杂度：N
空间复杂度：N
 */
func mirrorTree(root *common.TreeNode) *common.TreeNode {
	// 退出条件
	if root == nil{
		return root
	}

	// 交换左右节点
	root.Left,root.Right = root.Right,root.Left

	// 下沉
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root
}
