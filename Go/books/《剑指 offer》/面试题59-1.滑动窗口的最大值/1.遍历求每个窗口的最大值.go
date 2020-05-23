package main

func main()  {
	
}

/**
解题思路：遍历每个窗口，求每个窗口的最大值
时间复杂度：Nk
空间复杂度：N-k+1(窗口个数)
 */
func maxSlidingWindow(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	
	n := len(nums)
	
	var result []int
	for i:=0;i<n-k+1;i++{ // 遍历 n-k+1 个窗口
		var max = nums[i] // Notes: 初始化为窗口第一个元素
		// 构建窗口并求最大值
		for j:=i;j<i+k;j++{ 
			if nums[j] > max {
				max = nums[j]
			}
		}
		
		result = append(result,max)
	}

	return result
}	