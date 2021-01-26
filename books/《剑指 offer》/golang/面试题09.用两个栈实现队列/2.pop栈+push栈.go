package main

type CQueue2 struct {
	push []int 
	pop []int 
}


func Constructor2() CQueue2 {
	return CQueue2{}
}

/**
入队思路：直接放入 push 栈
时间复杂度：1
空间复杂度：N
*/
func (this *CQueue2) AppendTail(value int)  {
	this.push = append(this.push,value)
}

/**
出队思路：
1. 若 push 有值，则把 push 放入 pop
2. 若 push 无值，则直接从 pop 弹出最后一个元素
时间复杂度：N
空间复杂度：N
*/
func (this *CQueue2) DeleteHead() int {
	if len(this.pop) == 0 {
		if len(this.push) == 0 {
			return -1
		}
		// 把 push 放入 pop
		for i:= len(this.push)-1;i>=0;i--{
			this.pop = append(this.pop,this.push[i])
		}
		// 清空 push
		this.push = this.push[0:0]
	}
	// 从 pop 弹出
	result := this.pop[len(this.pop)-1]
	this.pop = this.pop[:len(this.pop)-1]
	
	return result
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */