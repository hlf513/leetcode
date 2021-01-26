package main

func lastRemaining2(n int, m int) int {
	var f = 0 // n=1,return 0
	// 逐个删除
	for i := 2; i != n+1; i++ { // n != n+1
		f = (m + f) % i
	}
	return f
}
