package main

import "fmt"

func main()  {
	
	fmt.Println(validateStackSequences([]int{1,2,3,4,5},[]int{4,5,3,2,1}))
	
}

/**
解题思路：pushed 压栈，判断栈顶元素和 popped 待出队的第一个元素是否相等，相等则从pushed 出栈，并从 popped 中删除
时间复杂度：N
空间复杂度：N (压栈需要新建一个栈)
*/
func validateStackSequences(pushed []int, popped []int) bool {
	l := len(pushed)
	j := 0 // 定义 popped 的待出栈下标
	var stack []int // 定义 pushed 需要的栈
	for i:=0;i<l;i++{
		stack = append(stack,pushed[i])		
		// fmt.Println("push:",pushed[i])
		for len(stack) > 0 && j < len(popped) && stack[len(stack)-1] == popped[j]{
			// fmt.Println("- pop:",popped[j])
			j++ // popped 下标后移
			stack = stack[:len(stack)-1] // 出栈
		}
	}
	
	return j == len(popped)

}