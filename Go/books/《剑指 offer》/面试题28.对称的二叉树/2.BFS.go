package main

import "leetcode/common"

/**
解题思路：BFS，层序遍历，每层一次弹出两个元素判断是否对称(值是否相等)
时间复杂度：N
空间复杂度：N
*/
func isSymmetric2(root *common.TreeNode) bool{
	// 边界条件
	if root == nil {
		return true
	}

	// 队列
	var q []*common.TreeNode
	// 注意，root 要入队两次；因为每次要取两个元素
	q = append(q,root) 
	q = append(q,root)

	for len(q) != 0 {
		// 出队两个元素
		left := q[0]
		right:=q[1]
		// 更新队列
		q = q[2:]

		// 开始判断是否对称
		if left == nil && right == nil{
			// 注意,不能直接返回 true，否则可能会有节点未判断
			// 示例：[9,-42,-42,null,76,76,null,null,13,null,13]
			continue
		}
		if left == nil || right == nil{
			return false
		}
		if left.Val != right.Val {
			return false
		}

		// 判断结束，把子节点写入队列，注意写入顺序
		q = append(q,left.Left)
		q = append(q,right.Right)
		q = append(q,right.Left)
		q = append(q,left.Right)
	}

	return true
}