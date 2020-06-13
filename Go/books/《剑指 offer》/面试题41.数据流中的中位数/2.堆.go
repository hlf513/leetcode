package main

/**
解题思路：大顶堆+小顶堆；各保存 n/2 元素，大顶堆保存前 n/2，小顶堆保存后 n/2，偶数 n中位数 = 大顶堆的堆顶；奇数 n 中位数 = 大小堆的堆顶。插入时判断小顶堆堆顶，>= 放入小顶堆，并判断个数是否满足
时间复杂度：插入 Ologn，查询 1
空间复杂度：N (两个堆保存所有数据)
*/
type MedianFinder2 struct {
	minArr []int
	maxArr []int
}

/** initialize your data structure here. */
func Constructor2() MedianFinder2 {
	heap := MedianFinder2{
		minArr: make([]int, 0),
		maxArr: make([]int, 0),
	}

	return heap
}

func (this *MedianFinder2) AddNum(num int) {
	// 第一个元素，扔到小根堆
	if len(this.minArr) == 0 {
		this.minHeapUp(num)
		return
	}

	// 元素比小根堆的堆顶元素大
	if num > this.minArr[0] {
		// 判断如果将该元素扔到小根堆时，如果此时小根堆长度比大根堆大1
		if len(this.minArr)+1-len(this.maxArr) > 1 {
			this.maxHeapUp(this.minArr[0]) // 将小根堆的堆顶元素扔到大根堆
			this.minHeapDown(num)          // 该元素放在小根堆的堆顶
		} else {
			this.minHeapUp(num) // 直接扔进小根堆
		}

	} else { // 元素比小根堆的堆顶元素小，需要扔到大根堆
		if len(this.maxArr)+1 > len(this.minArr) { // 如果将该元素扔到大根堆后，长度比小根堆要大，则不能直接扔进大根堆
			if num < this.maxArr[0] { // 该元素比大根堆的堆顶元素小
				this.minHeapUp(this.maxArr[0]) // 将大根堆的堆顶元素扔进小根堆
				this.maxHeapDown(num)          // 将该元素放在大根堆堆顶
			} else {
				this.minHeapUp(num) // 该元素比大根堆的堆顶元素大，直接扔到小根堆
			}
		} else {
			this.maxHeapUp(num) // 如果将该元素扔到大根堆后，长度不比小根堆大，则直接扔进大根堆
		}
	}
}

func (this *MedianFinder2) FindMedian() float64 {
	if len(this.minArr) == len(this.maxArr) {
		return (float64(this.minArr[0]) + float64(this.maxArr[0])) / 2
	} else {
		return float64(this.minArr[0])
	}
}

// 小根堆的上浮操作
func (this *MedianFinder2) minHeapUp(val int) {
	this.minArr = append(this.minArr, val)
	i := len(this.minArr) - 1
	for i > 0 && this.minArr[i] < this.minArr[(i-1)/2] {
		swap(this.minArr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// 大根堆的上浮操作
func (this *MedianFinder2) maxHeapUp(val int) {
	this.maxArr = append(this.maxArr, val)
	i := len(this.maxArr) - 1
	for i > 0 && this.maxArr[i] > this.maxArr[(i-1)/2] {
		swap(this.maxArr, i, (i-1)/2)
		i = (i - 1) / 2
	}
}

// 小根堆更新堆顶元素后的操作
func (this *MedianFinder2) minHeapDown(val int) {
	index := 0
	l := index*2 + 1

	minest := 0
	this.minArr[0] = val

	for l < len(this.minArr) {
		if l+1 < len(this.minArr) && this.minArr[l] > this.minArr[l+1] {
			minest = l + 1
		} else {
			minest = l
		}

		if this.minArr[index] < this.minArr[minest] {
			minest = index
		}

		if index == minest {
			break
		}

		swap(this.minArr, index, minest)
		index = minest
		l = index*2 + 1
	}
}

// 大根堆更新堆顶元素后的操作
func (this *MedianFinder2) maxHeapDown(val int) {
	index := 0
	l := index*2 + 1

	maxest := 0
	this.maxArr[0] = val

	for l < len(this.maxArr) {
		if l+1 < len(this.maxArr) && this.maxArr[l] < this.maxArr[l+1] {
			maxest = l + 1
		} else {
			maxest = l
		}

		if this.maxArr[index] > this.maxArr[maxest] {
			maxest = index
		}

		if index == maxest {
			break
		}

		swap(this.maxArr, index, maxest)
		index = maxest
		l = index*2 + 1
	}
}

// 交换两个数
func swap(heapArr []int, i, j int) {
	tmp := heapArr[i]
	heapArr[i] = heapArr[j]
	heapArr[j] = tmp
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
