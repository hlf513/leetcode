package main

/**
题目：返回中位数
解题思路：排序+返回中位下标
时间复杂度：排序算法
空间复杂度：N
*/
type MedianFinder struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	this.Nums = append(this.Nums, num)
}

func (this *MedianFinder) FindMedian() float64 {
	quickSort(this.Nums)
	l := len(this.Nums)
	if l%2 == 0 {
		return float64(this.Nums[l/2]+this.Nums[l/2-1]) / 2
	} else {
		return float64(this.Nums[(l-1)/2])
	}
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := nums[0]
	left, right := 0, len(nums)-1
	for i := 1; i <= right; {
		if nums[i] > mid {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		} else {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		}
	}
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
