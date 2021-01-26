package main

type MinStack struct {
	Main []int // 普通栈
	Mins []int // 保存最小元素的栈
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


/**
压栈思路：
1. 保存最小栈：压栈时比较最小栈顶数值，若<=则放入最小栈
2. 常规压栈
时间复杂度：1
空间复杂度：2N
*/
func (this *MinStack) Push(x int)  {
	if len(this.Main) == 0 || x <= this.Min() { // 注意，这个一定要是 <= 因为相同的最小值也要存储起来 
		this.Mins = append(this.Mins,x)
	}
	this.Main = append(this.Main,x)
	// fmt.Println("main",this.Main)
	// fmt.Println("min",this.Mins)
}


/**
出栈思路：
1. 维护最小栈：若出栈的数值是最小数值，则从最小栈出栈
2. 常规出栈
时间复杂度：1
空间复杂度：2N
 */
func (this *MinStack) Pop()  {
	if this.Top() == this.Min() {
		this.Mins = this.Mins[:len(this.Mins)-1]
	}
	this.Main = this.Main[:len(this.Main)-1]
}

/**
栈顶：直接查看最后一个元素
 */
func (this *MinStack) Top() int {
	return this.Main[len(this.Main)-1]
}

/**
最小值：直接查看最小栈的最后一个元素
 */
func (this *MinStack) Min() int {
	return this.Mins[len(this.Mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
