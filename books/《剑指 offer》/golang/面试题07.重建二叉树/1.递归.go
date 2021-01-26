package main

import "leetcode/common"

func main() {
	preorder := []int{2, 1, 3}
	inorder := []int{1, 2, 3}
	head := buildTree(preorder, inorder)
	common.TreeNode{}.PreOrder(head)
	//Common.TreeNode{}.InOrder(head)
	//Common.TreeNode{}.PostOrder(head)

}

/**
题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
解题思路：前序遍历中的第一个节点就是根节点，从中序遍历中找到根节点的下标，然后递归重建
时间复杂度：N
空间复杂度：N
*/
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	// 退出条件
	if len(preorder) == 0 {
		return nil
	}
	
	// 前序遍历中的第一个节点就是根节点
	var root = new(common.TreeNode)
	root.Val = preorder[0]

	// 在中序遍历里找到根节点的下标
	rootPos := 0
	for i := 0; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			rootPos = i
			break
		}
	}
	
	// 下沉
	root.Left = buildTree(preorder[1:rootPos+1], inorder[:rootPos])
	root.Right = buildTree(preorder[rootPos+1:], inorder[rootPos+1:])
	
	return root
}

