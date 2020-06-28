package main

/**
解题思路：维护一个单调栈，元素为 nums 下标（栈内元素对应的nums值递减，若入栈数大于栈顶数，则需要把小于入栈数都出栈）
时间复杂度：N
空间复杂度：N
*/
func maxSlidingWindow2(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	var stack,result []int
	// 遍历前 K 个元素，构建单调栈，找出第一个最大值
	for i:=0;i<k;i++{
		cleanStack(&stack,&nums,i,k)
	}
	result = append(result,nums[stack[0]])

	// 遍历其余元素，每次取单调栈的栈底元素作为最大值
	for i:=k;i<len(nums);i++{
		cleanStack(&stack,&nums,i,k)
		result = append(result,nums[stack[0]])
	}

	return result
}


func cleanStack(stack,nums *[]int, i ,k int)  {
	// 清理不是当前窗口的元素
	if len(*stack) > 0 && (*stack)[0] == i-k { // 判断第一个元素是不是需要清除
		*stack = (*stack)[1:]
	}
	// 保证栈是递减
	if len(*stack) > 0 {
		for j:=len(*stack) -1 ;j >=0;j--{
			if (*nums)[i] > (*nums)[(*stack)[j]]{
				*stack = (*stack)[:j]
			}
		}
	}
	// 压栈
	*stack = append(*stack,i)
}