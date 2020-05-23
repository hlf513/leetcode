package main

func main()  {
	
}

type CQueue struct {
	S1 []int // 队列栈
	S2 []int // 辅助栈
}


func Constructor() CQueue {
	return CQueue{}
}

/**
入队思路：
1. 若 S1 有值，则全部放入 S2 后，再放入 S1（改变出栈顺序）
2. 若 S1 无值，直接插入 S1
时间复杂度：2N
空间复杂度：2N
 */
func (this *CQueue) AppendTail(value int)  {
	l1 := len(this.S1)
	if l1 == 0 {
		this.S1 = append(this.S1,value)
	}else{
		// 从 S1 出栈 放入 S2
		for i:=l1-1;i>=0;i--{
			this.S2 = append(this.S2,this.S1[i])	
		}
		// 清空 S1
		this.S1 = []int{}
		// 放入插入的数据
		this.S1 = append(this.S1,value)
		// 从 S2出栈 放入 S1
		for i:=l1-1;i>=0;i--{
			this.S1 = append(this.S1,this.S2[i])
		}
		// 清空 S2
		this.S2 = []int{}
	}
}

/**
出队思路：
1. 若 S1 有值，则出栈
2. 若 S1 无值，返回-1
时间复杂度：1
空间复杂度：N
 */
func (this *CQueue) DeleteHead() int {
	l := len(this.S1)
	if l == 0{
		return -1
	}
	head := this.S1[len(this.S1)-1]
	this.S1 = this.S1[:len(this.S1)-1]
	return head
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */