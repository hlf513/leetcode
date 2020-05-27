package main

import "leetcode/common"

/**
题目：正序-倒序-正序打印
解题思路：BFS，层序遍历，每个层级追加入二维数组时，进行正倒序
时间复杂度：N
空间复杂度：N
 */
func levelOrder(root *common.TreeNode) [][]int {
	var result [][]int
	
	if root == nil{
		return result
	}
	
	// 初始化队列
	var q []*common.TreeNode
	q = append(q,root) // 根节点放入队列
	
	for len(q) != 0 {
		// 遍历每个层级
		var cur []int
		
		for _,item := range q{
			// 出队
			cur = append(cur,item.Val)
			q = q[1:]
			
			// 入队
			if item.Left != nil{
				q = append(q,item.Left)
			}
			if item.Right != nil{
				q = append(q,item.Right)
			}
		}

		// 保存当前层级结果
		if len(result) % 2 == 0 { // 第一个层级len=0（还没有写入 result），正序
			result = append(result,cur)
		}else{ // 倒序
			var tmp []int
			for i:= len(cur)-1;i>=0;i--{
				tmp = append(tmp,cur[i])
			}
			result = append(result,tmp)
		}
	}
	
	return result
}
