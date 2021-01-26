package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param2 := obj.Pop()
	param3 := obj.Peek()
	param4 := obj.Empty()

	fmt.Println(param2, param3, param4)
}

/**
用两个栈模拟
*/
type MyQueue struct {
	input  []int // 输入栈
	output []int // 输出栈
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var res int
	if len(this.output) > 0 {
		res = this.output[len(this.output)-1]
		this.output = this.output[:len(this.output)-1]
	} else {
		for i := len(this.input) - 1; i >= 0; i-- {
			// 从输入栈移动数据到输出栈
			item := this.input[i]
			this.input = this.input[:len(this.input)-1]
			this.output = append(this.output, item)
		}
		//fmt.Println(this.input, this.output)
		res = this.output[len(this.output)-1]
		this.output = this.output[:len(this.output)-1]
	}

	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	var res int
	if len(this.output) > 0 {
		res = this.output[len(this.output)-1]
	} else {
		for i := len(this.input) - 1; i >= 0; i-- {
			item := this.input[i]
			this.input = this.input[:len(this.input)-1]
			this.output = append(this.output, item)
		}
		//fmt.Println(this.input, this.output)
		res = this.output[len(this.output)-1]
	}

	return res
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.input) == 0 && len(this.output) == 0 {
		return true
	}
	return false
}
