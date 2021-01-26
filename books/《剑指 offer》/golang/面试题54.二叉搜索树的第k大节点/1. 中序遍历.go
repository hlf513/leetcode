package main

import (
	"leetcode/common"
)

func main()  {
	
}

/**
解题思路：利用二叉搜索树中序遍历是有序的特性，
时间复杂度：N
空间复杂度：N
 */
func kthLargest(root *common.TreeNode, k int) int {
	var ret []int

	// 中序遍历
	recurTree(root, &ret)

	// 获取第 K 大节点
	return ret[len(ret)-k]
}

func recurTree(root *common.TreeNode, ret *[]int)  {
	// 退出条件
	if root == nil{
		return 
	}
	
	if root.Left != nil{
		recurTree(root.Left,ret)		
	}
	
	*ret = append(*ret,root.Val)
	
	if root.Right != nil{
		recurTree(root.Right,ret)
	}
}
