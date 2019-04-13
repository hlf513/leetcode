package main

import "leetcode/Solutions/Go/Common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *Common.TreeNode) []int {
	var ret []int
	helper(root, &ret)

	return ret
}

func helper(root *Common.TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ret)

	*ret = append(*ret, root.Val)

	helper(root.Right, ret)
}
