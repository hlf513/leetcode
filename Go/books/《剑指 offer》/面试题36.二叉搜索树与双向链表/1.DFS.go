package main

import "leetcode/common"

func main() {

}

/**
题目：输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。要求不能创建任何新的节点，只能调整树中节点指针的指向。
解题思路：中序遍历，转为有序队列。
时间复杂度：
空间复杂度：
*/
var pre, head *common.TreeNode

func treeToDoublyList(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	dfs(root)
	// 把头节点连接到尾节点上，形成循环双向链表
	head.Left = pre
	head.Right = head

	return head
}

func dfs(root *common.TreeNode) {
	// 退出条件
	if root == nil {
		return
	}
	// 下沉
	if root.Left != nil {
		dfs(root.Left)
	}
	// logic
	if pre != nil {
		pre.Right = root // 前节点指向自己
	} else {
		head = root // 首次指向头节点
	}
	root.Left = pre // 当前节点指向前节点
	pre = root      // 用于存储前节点
	// 下沉
	if root.Right != nil {
		dfs(root.Right)
	}

}
