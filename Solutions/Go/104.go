package main

import (
	"fmt"
	"leetcode/Solutions/Go/Common"
)

func main() {
	t1 := Common.TreeNode{5, nil, nil}
	t2 := Common.TreeNode{1, nil, nil}
	t3 := Common.TreeNode{4, nil, nil}
	t4 := Common.TreeNode{}
	t5 := Common.TreeNode{}
	t6 := Common.TreeNode{3, nil, nil}
	t7 := Common.TreeNode{6, nil, nil}
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
func maxDepth(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	max, min := 1, 1

	dfs104(root, 1, &max, &min)

	//fmt.Println(max, min)

	return max
}

func dfs104(root *Common.TreeNode, level int, max *int, min *int) {
	if root == nil {
		return
	}
	if root.Right == nil && root.Left == nil {
		if *max < level {
			*max = level
		}
		if level < *min {
			*min = level
		}
	}

	dfs104(root.Left, level+1, max, min)
	dfs104(root.Right, level+1, max, min)
}

// 方法三
func maxDepth2(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*Common.TreeNode
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
