package main

import "strconv"

/**
解题思路： 按照字符串排序后拼接字符串，使用快排自定义排序规则
时间复杂度：logn
空间复杂度：N
 */
func minNumber2(nums []int) string {
	// int 转 string
	s := make([]string, len(nums))
	for i:=0; i<len(nums); i++ {
		s[i] = strconv.Itoa(nums[i])
	}
	// 排序
	quickSort(s)
	// 拼接
	ret := ""
	for _, v := range s {
		ret += v
	}
	return ret
}

func quickSort(data []string)  {
	// 退出条件
	if len(data) <= 1{
		return
	}
	
	// 寻找中位数
	mid := data[0]
	// 定义双指针
	head,tail := 0,len(data)-1
	
	// 从小到大排序
	for i:=1;i<=tail; {
		if data[i]+mid > mid +data[i] { // data[i] > mid
			// data[i] 放入右侧
			data[i],data[tail] = data[tail],data[i]
			// 右侧下标左移
			tail--
		}else{ // data[i] <= mid
			// data[i] 放入左侧
			data[i],data[head] = data[head],data[i]
			// 左侧下标右移
			head++
			// 
			i++
		}
	}
		
	quickSort(data[:head])
	quickSort(data[head+1:])
}

