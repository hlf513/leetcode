package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {

}

/**
解题思路：利用回溯思想，逐个复制，注意要保存已访问的节点，否则会出现死循环
时间复杂度：N
空间复杂度：N
*/
// 保存已复制节点
var visited = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	// 判断是否已复制
	if v, ok := visited[head]; ok {
		return v
	}
	// 新建节点进行复制
	n := &Node{Val: head.Val}
	// 标记当前节点已复制
	visited[head] = n
	// 开始复制
	n.Next = copyRandomList(head.Next)
	n.Random = copyRandomList(head.Random)

	// 返回头节点
	return n
}
