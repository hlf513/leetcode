package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

// DFS
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	helper22(&res, "", n, n)
	return res
}

func helper22(res *[]string, s string, left, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	if left > 0 {
		helper22(res, s+"(", left-1, right)
	}
	if right > left {
		helper22(res, s+")", left, right-1)
	}
}
