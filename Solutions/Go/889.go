package main

import (
	"leetcode/Solutions/Go/Common"
)

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	head := buildTree(preorder, postorder)
	Common.TreeNode{}.PreOrder(head)
	//Common.TreeNode{}.InOrder(head)
	//Common.TreeNode{}.PostOrder(head)

}

/**
方法：递归（根据遍历找规律，先找根，再找左右节点）
*/
func buildTree(preOrder []int, postOrder []int) *Common.TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	var root = new(Common.TreeNode)
	root.Val = preOrder[0]

	// 一个一个节点处理
	if len(preOrder) == 1 {
		return root
	}

	n := len(postOrder)
	for i := 0; i < n; i++ {
		if preOrder[1] == postOrder[i] {
			root.Left = buildTree(preOrder[1:i+2], postOrder[:i+1])
			root.Right = buildTree(preOrder[i+2:], postOrder[i+1:n-1])
		}

	}
	return root
}
