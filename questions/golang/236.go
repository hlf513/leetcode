package main

import (
	"fmt"
	"leetcode/common"
)

func main() {
	//[6,2,8,0,4,7,9,null,null,3,5]
	t1 := common.TreeNode{6, nil, nil}
	t2 := common.TreeNode{2, nil, nil}
	t3 := common.TreeNode{8, nil, nil}
	t4 := common.TreeNode{0, nil, nil}
	t5 := common.TreeNode{4, nil, nil}
	t6 := common.TreeNode{7, nil, nil}
	t7 := common.TreeNode{9, nil, nil}
	t8 := common.TreeNode{}
	t9 := common.TreeNode{}
	t10 := common.TreeNode{3, nil, nil}
	t11 := common.TreeNode{5, nil, nil}

	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t4
	t2.Right = &t5
	t3.Left = &t6
	t3.Right = &t7
	t4.Left = &t8
	t4.Right = &t9
	t5.Left = &t10
	t6.Right = &t11

	res := lowestcommonAncestor3(&t1, &t2, &t3)
	fmt.Println(res.Val)

}

/**
方法一：找出祖先路径（单链表），再计算两个单链表的相交节点
方法二：利用二叉树的特性
*/
// 方法二
func lowestcommonAncestor3(root, p, q *common.TreeNode) *common.TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	left := lowestcommonAncestor3(root.Left, p, q)
	right := lowestcommonAncestor3(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
