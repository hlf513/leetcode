package main

/**
题目：请实现一个函数用来匹配包含'. '和'*'的正则表达式；模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）
解题思路： 动态规划，画出二维数组，状态转移方程：
	- 字母：dp[i][j] = dp[i-1][j-1] // s[i] = p[j]
	- .: dp[i][j] = dp[i-1][j-1] // s[i] != ''
	- *:
		- 0次：dp[i][j] = dp[i][j-2] // p 消除「字母+*」组合
		- 多次：dp[i][j] = dp[i-1][j] ~ dp[i][j-2] // 匹配 s 末尾的字符，直到未匹配为止
时间复杂度：
空间复杂度：
*/
func isMatch(s string, p string) bool {
	m, n := len(s), len(p) // 计算s、p长度
	// matches 匹配前[i][j]
	matches := func(i, j int) bool {
		// 跳过 s 第一行（第一行表示空字符串）
		if i == 0 {
			return false
		}
		// 从s第二行（第一个字符）开始
		if p[j-1] == '.' { // 针对.匹配
			return true
		}
		return s[i-1] == p[j-1] // 针对字母进行匹配
	}

	// 初始化一个二维数组 [n+1][m+1]，表示「前[i][j]」是否匹配；初始化为 false
	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	// 双空字符串为 true
	f[0][0] = true
	for i := 0; i <= m; i++ { // s 字符串
		for j := 1; j <= n; j++ { // p 模式串
			// 开始检查前一个字符是否匹配
			if p[j-1] == '*' { // 若前一个字符是 *
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) { // (p）用*前字母进行匹配，
					f[i][j] = f[i][j] || f[i-1][j] // 若匹配则，再匹配s的前一个字符
				}
			} else if matches(i, j) { //（p）若前一个字符是.或者字母
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

func main() {
	isMatch("abc", "abc*")
}
