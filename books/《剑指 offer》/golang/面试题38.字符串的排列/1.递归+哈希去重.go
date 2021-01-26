package main

/**
题目：输入一个字符串，打印出该字符串中字符的所有排列。
解题思路：递归+哈希去重，DFS，每次固定一位，其他位递归前交换，递归后还原
时间复杂度：N!
空间复杂度：1
*/
func permutation(s string) []string {
	var res []string
	helper([]byte(s), 0, &res)
	return res
}

// helper 固定每位，再排列其他位
// s:排列后的字符串
// start:递归参数，字符串下标
// res:保存排列后的结果
func helper(s []byte, start int, res *[]string) {
	// 遍历到最后一位，则存储字符串
	if start == len(s) {
		*res = append(*res, string(s))
	}
	// m 用于当前轮次的字符判重
	m := make(map[byte]int)
	// 遍历每一位
	for i := start; i < len(s); i++ {
		// 重复字符
		if _, ok := m[s[i]]; ok {
			continue
		}
		// 依次和 start 下标交换
		s[i], s[start] = s[start], s[i]
		// 下沉
		helper(s, start+1, res)
		// 还原交换
		s[i], s[start] = s[start], s[i]
		// 记录字符
		m[s[i]] = 1
	}
}
