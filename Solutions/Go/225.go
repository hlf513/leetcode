package main

import (
	"fmt"
)

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	//obj.Push(3)
	//param_4 := obj.Top()
	param_2 := obj.Pop()
	param_4 := obj.Empty()

	fmt.Println(param_2, param_3, param_4)
}

type MyStack struct {
	q []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	//fmt.Println("pushed:",this.q)
	if len(this.q) > 1 {
		// 1234 => 12341
		for i := 1; i < len(this.q); i++ {
			this.q = append(this.q, this.q[0])
			this.q = this.q[1:len(this.q)]
		}
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.q[0]
	this.q = this.q[1:len(this.q)]
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.q) == 0 {
		return true
	}
	return false
}
