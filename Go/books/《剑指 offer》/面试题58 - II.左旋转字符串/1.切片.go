package main

/**
解题思路：利用切换/字符数组
时间复杂度：N
空间复杂度：1
*/
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
