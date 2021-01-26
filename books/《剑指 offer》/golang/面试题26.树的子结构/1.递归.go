package main

import "leetcode/common"

func main() {

}

/**
题目：输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
解题思路：两个递归，helper()判断a是否包含b，isSubStructure()判断A每个节点作为根节点时调用helper()判断是否包含B
时间复杂度：MN
空间复杂度：M(A的节点数量)
*/
func isSubStructure(A *common.TreeNode, B *common.TreeNode) bool {
	// 退出条件
	if A == nil || B == nil {
		return false
	}

	// 下沉
	return helper(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 判断 A 是否包含 B
func helper(A *common.TreeNode, B *common.TreeNode) bool {
	// 退出条件
	if B == nil { // A有节点，B无节点
		return true
	}
	if A == nil { // B有节点，A有节点
		return false
	}

	return A.Val == B.Val && // AB节点相等
		helper(A.Left, B.Left) && // 比较AB左节点
		helper(A.Right, B.Right) // 比较AB右节点
}
