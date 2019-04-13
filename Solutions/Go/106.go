package main

import (
	"leetcode/Solutions/Go/Common"
)

func main() {
	inOrder := []int{1, 2, 3}
	postOrder := []int{1, 3, 2}
	head := buildTree(inOrder, postOrder)
	Common.TreeNode{}.PreOrder(head)
	//Common.TreeNode{}.InOrder(head)
	//Common.TreeNode{}.PostOrder(head)
}

/**
方法：递归（根据遍历找规律，先找根，再找左右子树）
*/
func buildTree(inOrder []int, postOrder []int) *Common.TreeNode {
	if len(inOrder) == 0 {
		return nil
	}
	root := new(Common.TreeNode)
	root.Val = postOrder[len(postOrder)-1]

	for i := 0; i < len(inOrder); i++ {
		if root.Val == inOrder[i] {
			root.Left = buildTree(inOrder[:i], postOrder[:i])
			root.Right = buildTree(inOrder[i+1:], postOrder[i:len(postOrder)-1])
		}
	}

	return root
}
