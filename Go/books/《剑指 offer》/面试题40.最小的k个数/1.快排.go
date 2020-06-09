package main

/**
解题思路： 排序后，取前 k 值
时间复杂度：NlogN（取决与排序算法）
空间复杂度：N(k 个值组成的数组）
*/
func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr)
	return arr[:k]
}

func quickSort(data []int) {
	// 退出条件
	if len(data) <= 1 {
		return
	}
	// 处理逻辑
	mid := data[0]               // 待比较数据
	head, tail := 0, len(data)-1 // 下标
	for i := 1; i <= tail; {
		if data[i] > mid { // > 表示从小到大；< 表示从大到小
			// 与尾部交换，重新判断
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			// 与头部交换，移动下标
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	// 下沉
	quickSort(data[:head])
	quickSort(data[head+1:])
}
