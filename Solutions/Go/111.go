package main

import (
	"fmt"
	"leetcode/Solutions/Go/Common"
)

func main() {
	//  5
	// 1,4
	// 3,6
	t1 := Common.TreeNode{5, nil, nil}
	t2 := Common.TreeNode{1, nil, nil}
	t3 := Common.TreeNode{4, nil, nil}
	t6 := Common.TreeNode{3, nil, nil}
	t7 := Common.TreeNode{6, nil, nil}
	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t6
	t2.Right = &t7

	fmt.Println(minDepth2(&t1))
}

/**
方法一：DFS（判断是不是叶子节点，变量：max、min）
方法二：BFS（层级）
*/

// 方法一
func minDepth(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	var min = 0
	dfs111(root, 1, &min)

	return min
}

func dfs111(root *Common.TreeNode, level int, min *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if *min == 0 {
			*min = level
		} else if level < *min {
			*min = level
		}
	}

	dfs111(root.Left, level+1, min)
	dfs111(root.Right, level+1, min)
}

// 方法二
func minDepth2(root *Common.TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*Common.TreeNode
	q = append(q, root)
	level := 1

	for len(q) > 0 {
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return level
			}
			q = q[1:]

			// 注意：这里是 node 不是 root
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
