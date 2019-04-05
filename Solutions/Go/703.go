package main

import (
	"container/heap"
	"fmt"
)

func main() {
	k := 2
	nums := []int{3, 2, 5, 4}
	obj := Constructor(k, nums)
	param_1 := obj.Add(-1)
	param_2 := obj.Add(1)
	param_3 := obj.Add(-2)
	param_4 := obj.Add(-4)
	param_5 := obj.Add(3)

	fmt.Println(param_1, param_2, param_3, param_4, param_5)
}

type minHeap []int

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }
func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}
func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}

type KthLargest struct {
	mh *minHeap // 小顶堆，使用快排超时了...
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	h := &minHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.mh, val)
	if this.mh.Len() > this.k {
		heap.Pop(this.mh)
	}

	mh := *this.mh
	return mh[0]
}
