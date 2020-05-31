package main

import (
	"strconv"
	"strings"

	"leetcode/common"
)

// Constructor 初始化
func Constructor2() Codec {
	return Codec{}
}

/**
解题思路：BFS，队列
时间复杂度：N
空间复杂度：N
 */
func (this *Codec) serialize2(root *common.TreeNode) string {
	// 新建队列
	queue := []*common.TreeNode{root}
	vals := make([]string, 0)
	
	// 遍历	
	for len(queue) > 0 {
		// 出队
		node := queue[0] 
		queue = queue[1:]
		if node == nil {
			vals = append(vals, "null")
		} else {
			vals = append(vals, strconv.Itoa(node.Val))
			// 把左右子树入队
			queue = append(queue, node.Left, node.Right)
		}
	}
	
	// 格式化
	return strings.Join(vals, ",")
}


func (this *Codec) deserialize2(data string) *common.TreeNode {
	// 字符串切分
	vals := strings.Split(data, ",")
	// 边界条件
	idx := 0
	if vals[idx] == "null" { 
		return nil
	}
	// 初始化头节点	
	val, _ := strconv.Atoi(vals[idx])
	root := &common.TreeNode{Val: val}
	// 新建队列并入队
	queue := []*common.TreeNode{root}
	var node, left, right *common.TreeNode
	// 遍历
	for len(queue) > 0 {
		// 出队
		node = queue[0]
		queue = queue[1:]
		// 处理左子树
		idx++
		if vals[idx] != "null" {
			val, _ = strconv.Atoi(vals[idx])
			left = &common.TreeNode{Val: val}
			node.Left = left
			// 入队
			queue = append(queue, left)
		}
		// 处理右子树
		idx++
		if vals[idx] != "null" {
			val, _ = strconv.Atoi(vals[idx])
			right = &common.TreeNode{Val: val}
			node.Right = right
			// 入队
			queue = append(queue, right)
		}
	}
	// 返回头节点
	return root
}


