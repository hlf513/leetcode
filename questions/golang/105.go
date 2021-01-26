package main

import (
	"leetcode/common"
)

func main() {
	preorder := []int{2, 1, 3}
	inorder := []int{1, 2, 3}
	head := buildTree(preorder, inorder)
	common.TreeNode{}.PreOrder(head)
	//common.TreeNode{}.InOrder(head)
	//common.TreeNode{}.PostOrder(head)

}

/**
方法：递归（根据遍历找规律，先找根，再找左右子树）
*/
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//root := common.TreeNode{preorder[0], nil, nil}
	//var root *common.TreeNode
	var root = new(common.TreeNode)
	root.Val = preorder[0]

	rootPos := 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			rootPos = i
			break
		}
	}
	//rootPos := binarySearch(inorder, preorder[0])
	root.Left = buildTree(preorder[1:rootPos+1], inorder[:rootPos])
	root.Right = buildTree(preorder[rootPos+1:], inorder[rootPos+1:])
	return root
}
