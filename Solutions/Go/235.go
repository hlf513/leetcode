package main

import (
	"fmt"
	"leetcode/Solutions/Go/Common"
)

func main() {
	//[6,2,8,0,4,7,9,null,null,3,5]
	t1 := Common.TreeNode{6, nil, nil}
	t2 := Common.TreeNode{2, nil, nil}
	t3 := Common.TreeNode{8, nil, nil}
	t4 := Common.TreeNode{0, nil, nil}
	t5 := Common.TreeNode{4, nil, nil}
	t6 := Common.TreeNode{7, nil, nil}
	t7 := Common.TreeNode{9, nil, nil}
	t8 := Common.TreeNode{}
	t9 := Common.TreeNode{}
	t10 := Common.TreeNode{3, nil, nil}
	t11 := Common.TreeNode{5, nil, nil}

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

	res := lowestCommonAncestor2(&t1, &t2, &t3)
	fmt.Println(res.Val)

}

/**
利用平衡二叉树的特性进行判断
*/

// 递归
func lowestCommonAncestor(root, p, q *Common.TreeNode) *Common.TreeNode {
	// pq 在左子树
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	// pq 在右子树
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	// pq 在左右子树
	return root
}

// 非递归
func lowestCommonAncestor2(root, p, q *Common.TreeNode) *Common.TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
