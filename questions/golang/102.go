package main

import (
	"fmt"
	"leetcode/common"
)

func main() {
	//[5,1,4,null,null,3,6]
	t1 := common.TreeNode{5, nil, nil}
	t2 := common.TreeNode{1, nil, nil}
	t3 := common.TreeNode{4, nil, nil}
	t4 := common.TreeNode{}
	t5 := common.TreeNode{}
	t6 := common.TreeNode{3, nil, nil}
	t7 := common.TreeNode{6, nil, nil}
	t1.Left = &t2
	t1.Right = &t3
	t2.Left = &t4
	t2.Right = &t5
	t3.Left = &t6
	t3.Right = &t7

	fmt.Println(levelOrder2(&t1))
	// [[5] [1 4] [0 0 3 6]]
}

/**
方法一：BFS O(N)
方法二：DFS O(N)
*/

// 方法一
func levelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var q []*common.TreeNode
	var res [][]int
	q = append(q, root)

	for len(q) > 0 {
		var cur []int

		for _, node := range q { // 循环体内的 append 不影响这里的 range
			cur = append(cur, node.Val)
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, cur)
	}

	return res
}

// 方法二
func levelOrder2(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	m := make(map[int][]int)
	dfs(root, 0, &m)

	var res [][]int
	for i := 0; i < len(m); i++ {
		res = append(res, m[i])
	}

	return res
}

func dfs(node *common.TreeNode, level int, res *map[int][]int) {
	if node == nil {
		return
	}

	tmp := *res
	tmp[level] = append(tmp[level], node.Val)

	dfs(node.Left, level+1, &tmp)
	dfs(node.Right, level+1, &tmp)
}
