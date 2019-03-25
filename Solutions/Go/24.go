package main

func main() {
	tail := ListNode{5, nil}
	next3 := ListNode{4, &tail}
	next2 := ListNode{3, &next3}
	next1 := ListNode{2, &next2}
	head := ListNode{1, &next1}

	printListNode(*swapPairs(&head))
}

func swapPairs(head *ListNode) *ListNode {

}
