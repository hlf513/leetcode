package common

import "fmt"

// TreeNode 树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序遍历
func (t TreeNode) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	t.PreOrder(node.Left)
	t.PreOrder(node.Right)
}

// InOrder 中序遍历
func (t TreeNode) InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	t.InOrder(node.Left)
	fmt.Println(node.Val)
	t.InOrder(node.Right)
}

// PostOrder 后序遍历
func (t TreeNode) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	t.PostOrder(node.Left)
	t.PostOrder(node.Right)
	fmt.Println(node.Val)
}
