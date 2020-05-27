package main

import "leetcode/common"

/**
解题思路：广度优先搜索 BFS（队列思维），层序遍历，把每个父节点放入队列，然后交换子节点。
时间复杂度：N
空间复杂度：N
*/
func mirrorTree2(root *common.TreeNode) *common.TreeNode {
	// 退出条件
	if root == nil{
		return root
	}
	// 初始化队列
	var q []*common.TreeNode
	q = append(q,root)

	for len(q) != 0 {
		// 弹出头节点
		current := q[0]
		q=q[1:]
		//  交换
		current.Left,current.Right = current.Right,current.Left
		// 有子树则放入队列中，继续交换节点
		if current.Left != nil{
			q = append(q,current.Left)
		}
		if current.Right != nil{
			q = append(q,current.Right)
		}
	}

	return root
}