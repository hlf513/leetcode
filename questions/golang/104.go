package main

import (
	"fmt"

	"leetcode/common"
)

func main() {
	t1 := common.TreeNode{Val: 5}
	t2 := common.TreeNode{Val: 1}
	t3 := common.TreeNode{Val: 4}
	t4 := common.TreeNode{}
	t5 := common.TreeNode{}
	t6 := common.TreeNode{Val: 3}
	t7 := common.TreeNode{Val: 6}
	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t4
	t2.Right = &t5
	t3.Left = &t6
	t3.Right = &t7

	fmt.Println(maxDepth2(&t1))
}

/**
方法一：暴力解法（左右子树深度比大小）
方法二：DFS（判断是不是叶子节点，变量：max、min）
方法三：BFS（层级）
*/

// 方法二
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	max, min := 1, 0

	dfs104(root, 1, &max, &min)

	// fmt.Println(max, min)

	return max
}

func dfs104(root *common.TreeNode, level int, max *int, min *int) {
	if root == nil {
		return
	}
	if root.Right == nil && root.Left == nil {
		if *max < level {
			*max = level
		}
		if *min == 0 {
			*min = level
		} else if level < *min {
			*min = level
		}
	}

	dfs104(root.Left, level+1, max, min)
	dfs104(root.Right, level+1, max, min)
}

// 方法三
func maxDepth2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*common.TreeNode
	q = append(q, root)

	var level int

	for len(q) > 0 {
		for _, node := range q {
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}

	return level
}
