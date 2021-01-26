package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this TreeNode) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	this.PreOrder(node.Left)
	this.PreOrder(node.Right)
}

func (this TreeNode) InOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.InOrder(node.Left)
	fmt.Println(node.Val)
	this.InOrder(node.Right)
}

func (this TreeNode) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}
	this.PostOrder(node.Left)
	this.PostOrder(node.Right)
	fmt.Println(node.Val)
}
