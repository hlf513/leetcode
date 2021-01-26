package main

import (
	"strconv"
	"strings"

	"leetcode/common"
)

func main()  {
	
	
}

type Codec struct {
}

// Constructor 初始化
/**
解题思路：根据前序遍历(DFS)，递归序列化，反序列化
时间复杂度：N
空间复杂度：N
 */
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *common.TreeNode) string {
	if root == nil {
		return "null,"
	}
	
	// 保存顺序：根结点、左节点、右节点（前序遍历）
	return strconv.Itoa(root.Val) +"," +  this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *common.TreeNode {
	list := strings.Split(data,",")
	
	return dfs(list)
}

func dfs(values []string) *common.TreeNode {
	// 获取第一个元素
	val := values[0] 
	values = values[1:]
	
	// 退出条件
	if val == "null"{
		return nil
	}
	
	// 创建节点
	v,_ := strconv.Atoi(val)
	node := &common.TreeNode{Val:v}
	node.Left = dfs(values)
	node.Right = dfs(values)
	
	return node
}


/**
 * Your Codec object will b
e instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */