package main

import (
	"fmt"
	"leetcode/common"
	"math"
)

func main() {
	//[5,1,4,null,null,3,6]
	t1 := common.TreeNode{5, nil, nil}
	t2 := common.TreeNode{1, nil, nil}
	t3 := common.TreeNode{4, nil, nil}
	t4 := common.TreeNode{}
	t5 := common.TreeNode{}
	t6 := common.TreeNode{3, nil, nil}
	t7 := common.TreeNode{6, nil, nil}
	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t4
	t2.Right = &t5
	t3.Left = &t6
	t3.Right = &t7

	fmt.Println(isValidBST(&t1))
}

func isValidBST(root *common.TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

/**
思路：左孩子小于根节点，右孩子大于根节点
*/
func isValid(root *common.TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}
