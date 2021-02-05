package main

import "fmt"

func main() {
	tail := ListNode{5, nil}
	next3 := ListNode{4, &tail}
	next2 := ListNode{3, &next3}
	next1 := ListNode{2, &next2}
	head := ListNode{1, &next1}

	printListNode(*swapPairs(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head ListNode) {
	for {
		fmt.Println(head.Val)
		if head.Next == nil {
			break
		}
		head = *head.Next
	}
}

/**
 * 解题思路：虚拟节点
 * 		1. 声明 dummy 节点
 * 		2. 定义3个指针
 * 		3. 交换 cur 与 post，并移动 pre
 *		4. 循环判断 pre.next 与 pre.next.next 是否存在
 * 时间复杂度：O(?)
 * 空间复杂度：O(?)
*/
func swapPairs(head *ListNode) *ListNode {
	dummy := ListNode{}
	pre := &dummy
	pre.Next = head

	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		post := cur.Next

		pre.Next, post.Next, cur.Next = post, cur, post.Next

		pre = cur
	}

	return dummy.Next
}
